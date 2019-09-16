# -*- coding: utf-8 -*-
import scrapy
from scrapy.http import Request


class BaiduSpider(scrapy.Spider):
    name = 'baidu'
    allowed_domains = ['www.baidu.com']
    start_urls = ['http://www.baidu.com/']

    def parse(self, response):
        print(response.text)
        print('===================================')
        url = 'https://www.so.com/'
        return Request(url,callback = get_baidu_info)

    def get_baidu_info(self,response):
        print(response.text)
        print('=====================')

