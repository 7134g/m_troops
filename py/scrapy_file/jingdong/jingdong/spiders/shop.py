# -*- coding: utf-8 -*-
import scrapy
from scrapy import Request
from urllib import parse
from scrapy.loader import ItemLoader
from jingdong.items import ShopItem


class ShopSpider(scrapy.Spider):
    name = 'shop'
    allowed_domains = ['www.jd.com']
    start_urls = ['http://www.jd.com/']


    def start_requests(self):
        return [Request('登录页',callback=self.login)]

    def login(self,response):
        # 解析_xsrf参数值
        #
        #
        #
        # 得到
        xsrf = ''

        if xsrf:
            post_url = 'post页'
            post_data = {
            "_xsrf":xsrf,
            "account":'account',
            "password":'password'
            }

            return [scrapy.FormRequest(
                url = post_url,
                formdata = post_data,
                callback=self.check_login
            )]

    def check_login(self,response):
        # 验证服务器返回是否成功
        if response.status == 200:
            return True
        for url in self.start_urls:
            yield Request(url, dont_filter=True)


    def parse(self, response):
        all_url = response.css('').extract()
        all_url = [parse.urljoin(response.url,url) for url in all_url]
        all_url = filter(lambda x:True if x.startswith("https") else False,all_url)
        for url in all_url:
            if url:
                yield Request(url,callback=self.parse_question)
            pass

    def parse_question(self,response):
        Item_loader = ItemLoader(item = ShopItem(),response=response)
        item_loader.add_value('url','匹配式')
        item_loader.add_css('Id', '匹配式')
        item_loader.add_css('name', '匹配式')
        item_loader.add_css('title', '匹配式')

        response_item = item_loader.load_item()
        yield response_item

