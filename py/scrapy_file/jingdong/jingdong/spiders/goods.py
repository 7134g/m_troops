# -*- coding: utf-8 -*-
import scrapy
from scrapy.http import Request
import re
from jingdong.items import JingdongItem
from jingdong.settings import *
import requests
from fake_useragent import UserAgent
from scrapy_redis.spiders import RedisSpider


class GoodsSpider(RedisSpider):
    name = 'goods'
    # allowed_domains = ['jd.com']
    # start_urls = ['http://jd.com/']

    redis_key = "GoodsSpider:start_urls"

    # scrapy_redis动态域
    def __init__(self, *args, **kwargs):
        domain = kwargs.pop('domain', '')
        self.allowed_domains = filter(None, domain.split(','))
        super(GoodsSpider, self).__init__(*args, **kwargs)

    ua = UserAgent()
    headers = {
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
        'Accept-Language': 'zh-CN,zh;q=0.9',
        'user-agent': ua.random,
        'Connection': 'keep - alive',
        'referer': 'https://search.jd.com/Search?keyword=%E5%9B%BE%E4%B9%A6&enc=utf-8&wq=%E5%9B%BE%E4%B9%A6&page=1'

    }
    # https://search.jd.com/Search?keyword=%E5%9B%BE%E4%B9%A6&enc=utf-8&wq=%E5%9B%BE%E4%B9%A6&page=1
    # https://search.jd.com/Search?keyword=图书&enc=utf-8&wq=图书&page=1

    # 搜索的起始页
    url = "https://search.jd.com/Search?keyword={KEYWORDS}&enc=utf-8&wq={KEYWORDS}&page={page}"

    # 电子价格
    Eprice_url = "https://c.3.cn/book?skuId={skuId}&cat={cat}&area=1_72_2799_0&callback=book_jsonp_callback"

    # 商品价格
    price_url = "https://p.3.cn/prices/mgets?type=1&area=1_72_2799_0&pdtk=&pduid=1771569446&pdpin=&pdbp=0&skuIds=J_{skuId}&ext=11100000&callback=jQuery3021180&_=1547383556702"
    price2_url = 'https://c0.3.cn/stock?skuId={skuId}&venderId=1000005720&cat={cat}&area=1_72_2799_0&buyNum=1&extraParam={%22originid%22:%221%22}&ch=1&pduid=1771569446&pdpin=&fqsp=0&callback=getStockCallback'

    # 评论
    comment_url = "https://sclub.jd.com/comment/productPageComments.action?callback=fetchJSON_comment98vv39228&productId={skuId}&score=0&sortType=5&page=0&pageSize=10&isShadowSku=0&fold=1"

    def start_requests(self):
        for k in range(1, PAGE_NUM):
            yield Request(url=self.url.format(KEYWORDS=KEYWORDS, page=2 * k - 1), callback=self.page_parse)

    def page_parse(self, response):
        # 每页商品ID
        goodsID = response.xpath('//li/@data-sku').extract()
        print(goodsID)

        for each in goodsID:
            goodsurl = "https://item.jd.com/{}.html".format(each)
            yield Request(url=goodsurl, callback=self.get_goods_info)

    def get_goods_info(self, response):

        item = JingdongItem()

        # 图书链接
        item["link"] = response.url

        # 图书标题
        item["title"] = response.xpath('//div[@class="sku-name"]/text()').extract()[0].strip()

        # 作者
        item["writer"] = re.sub(' ', '', ''.join(response.xpath('//div[@class="p-author"]/a/text()').extract()))

        # 提取商品ID
        skuId = re.compile(r'https:..item.jd.com.(\d+).html').findall(response.url)[0]
        item['Id'] = skuId
        cat = re.compile(r'pcat:\[(.*?)\],').findall(response.text)
        cat = re.sub("\|", ",", cat[0]).strip("'")
        item['catId'] = cat
        print(skuId)
        print(cat)

        try:
            # 打开商品价格
            res_no_price = requests.get(url=self.price_url.format(skuId=skuId), headers=self.headers)
            item["n_price"] = re.compile('"op":"(.*?)",').findall(res_no_price.text)[0]
            item["o_price"] = re.compile('"m":"(.*?)",').findall(res_no_price.text)[0]
            # 打开电子书价格
            res_e_price = requests.get(url=self.Eprice_url.format(skuId=skuId, cat=cat), headers=self.headers)
            item["e_price"] = re.compile('"p":"(.*?)",').findall(res_e_price.text)[0]
        except IndexError:
            res_no_price = requests.get(url=self.price2_url.format(skuId=skuId, cat=cat), headers=self.headers)
            item["n_price"] = re.compile('"op":"(.*?)",').findall(res_no_price.text)[0]
            item["o_price"] = re.compile('"m":"(.*?)",').findall(res_no_price.text)[0]
        finally:
            # 用户评论
            res_comment = requests.get(url=self.comment_url.format(skuId=skuId), headers=self.headers)
            item["comment"] = re.compile('"content":"(.*?)",').findall(res_comment.text)

            yield item

