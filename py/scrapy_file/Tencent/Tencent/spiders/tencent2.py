# -*- coding: utf-8 -*-
import scrapy
from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import CrawlSpider, Rule
from Tencent.items import TencentItem


class Tencent2Spider(CrawlSpider):
    name = 'tencent2'
    allowed_domains = ['tencent.com']
    start_urls = ['https://hr.tencent.com/position.php?&start=0']

    rules = (
        Rule(LinkExtractor(allow="start=\d+"), callback='parse_item', follow=True),

    )

    def parse_item(self,response):
        # i = {}
        # #i['domain_id'] = response.xpath('//input[@id="sid"]/@value').extract()
        # #i['name'] = response.xpath('//div[@id="name"]').extract()
        # #i['description'] = response.xpath('//div[@id="description"]').extract()
        # return i

        item = TencentItem()
        node_list = response.xpath('//tr[@class="even"] | //tr[@class="odd"]')

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