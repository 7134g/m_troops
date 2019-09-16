# -*- coding: utf-8 -*-
import scrapy
import json
from Douyu.items import DouyuItem

class DouyuSpider(scrapy.Spider):
    name = 'douyu'
    allowed_domains = ['douyucdn.cn']

    baseURL = 'http://capi.douyucdn.cn/api/v1/getVerticalRoom?limit=20&offset='
    offset = 0
    start_urls = [baseURL+str(offset)]

    def parse(self, response):
        data_list = json.loads(response.body)['data']

        # 长度为0，数据已提取完毕
        if not len(data_list):
            return 

        for data in data_list:
            item = DouyuItem()

            # 获取数据
            item['nickname'] = data['nickname']
            item['vertical_src'] = data['vertical_src']

            yield item

        
            self.offset += 20
            url = self.baseURL + str(self.offset)
            yield scrapy.Request(url,callback = self.parse)