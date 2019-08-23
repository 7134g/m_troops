# -*- coding: utf-8 -*-
import scrapy
from Tencent.items import TencentItem


class TencentSpider(scrapy.Spider):
    # 爬虫名
    name = 'tencent'
    allowed_domains = ['tencent.com']
    baseURL = "https://hr.tencent.com/position.php?&start="
    offset = 0
    #拼接url
    start_urls = [baseURL + str(offset)]

    def parse(self, response):
        # 提取全部response数据
        node_list = response.xpath('//tr[@class="even"]|//tr[@class="odd"]')

        for node in node_list:
            item = TencentItem()


            # 提取每个职位信息，并将Unicode字符串编码成utf8
            item['position_name'] = node.xpath('./td[1]/a/text()').extract()[0]
            item['position_link'] = node.xpath('./td/a/@href').extract()[0]

            if len(node.xpath('./td[2]/text()')):
                item['position_type'] = node.xpath('./td[2]/text()').extract()[0]
            else:
                item['position_type'] = ''

            item['people_number'] = node.xpath('./td[3]/text()').extract()[0]
            item['work_place'] = node.xpath('./td[4]/text()').extract()[0]
            item['release_time'] = node.xpath('./td[5]/text()').extract()[0]

            yield item

        # 第一种：适用页面找不到链接，必须通过拼接获取
        # if self.offset < 2190:
        #     self.offset += 10
        #     url = self.baseURL + str(self.offset)
        #     yield scrapy.Request(url,callback = self.parse)

        # 第二种：页面可以获取url，遍历全部页
        # if not len(response.xpath("//a[@class='noactive' and @id='next']")):
        #     url = response.xpath("//a[@id='next']/@href").extract()[0]
        #     yield scrapy.Request("https://hr.tencent.com/" + url,callback = self.parse)

