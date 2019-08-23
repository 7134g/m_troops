# -*- coding: utf-8 -*-
import scrapy
from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import CrawlSpider, Rule
from dg_sunny.items import DgSunnyItem
import re


class MsgeSpider(CrawlSpider):
    name = 'msge'
    allowed_domains = ['wz.sun0769.com']
    start_urls = ['http://wz.sun0769.com/index.php/question/reply?page=0']

    rules = (
        Rule(LinkExtractor(allow='reply\?page=\d+'),follow=True),
        Rule(LinkExtractor(allow='http://.*\d+\.shtml'),callback = "detali_parse",follow = False)
    )

    def detali_parse(self, response):

        item = DgSunnyItem()

        item['url'] = response.url
        item['title'] = re.sub('\s','',response.css('.wzy1 .niae2_top::text').extract()[0])
        if not item['title']:
            item['title'] = re.sub('\s','',response.css('strong.tgray14::text').extract()[0])
        item['num'] = response.css('.wzy1 span[style]::text').re('\d+')[0]
        item['a_content'] = re.sub('\s','',response.css('.wzy1 .txt16_3::text').extract()[0])
        item['g_content'] = re.sub('\s','',''.join(response.css('.wzy1 .txt16_3::text').extract()[2:-1]))

        yield item