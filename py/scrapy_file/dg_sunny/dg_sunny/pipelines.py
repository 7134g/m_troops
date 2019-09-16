# -*- coding: utf-8 -*-
import json
from scrapy.exceptions import DropItem
# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html


class DgSunnyPipeline(object):
    def __init__(self):
        self.file = open('msge.json','w',encoding='utf-8')
        self.number = []

    def process_item(self, item, spider):
        # if item['url']:
        #     if item['num'] in self.number:
        #         raise DropItem('数据已存在')
        #     self.number.append(item['num'])
        text = json.dumps(dict(item),ensure_ascii= False)+',\n'
        self.file.write(text)
        return item

    def close_parse(self):
        self.file.close()