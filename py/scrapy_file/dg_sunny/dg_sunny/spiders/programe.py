# -*- coding: utf-8 -*-
import scrapy
from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import CrawlSpider, Rule
from dg_sunny.items import ProgrameItemLoader,ProgrameItem




class MsgeSpider(CrawlSpider):
    name = 'programe'
    allowed_domains = ['wz.sun0769.com']
    start_urls = ['http://wz.sun0769.com/index.php/question/reply?page=0']

    rules = (
        Rule(LinkExtractor(allow='reply\?page=\d+'),follow=True),
        Rule(LinkExtractor(allow='http://.*\d+\.shtml'),callback = "detali_parse",follow = False)
    )

    def detali_parse(self, response):

        item_loader = ProgrameItemLoader(item=ProgrameItem(),response=response)

        item_loader.add_value('url',response.url)
        try:
            item_loader.add_css('title','.wzy1 .niae2_top::text')
        except:
            item_loader.add_css('title','strong.tgray14::text')
        item_loader.add_css('num','.wzy1 span[style]::text')
        item_loader.add_css('a_content','.wzy1 .txt16_3::text')
        item_loader.add_css('g_content','.wzy1 .txt16_3::text')

        sun_item = item_loader.load_item()

        yield sun_item