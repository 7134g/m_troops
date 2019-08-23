# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html

from scrapy.exceptions import DropItem
import json
import os
from jingdong.settings import *
import pymongo


class JingdongPipeline(object):
    def __init__(self):
        self.f = open(os.getcwd()+'\data\jingdong.json','w')
        self.link_name = []

    # 要么返回一个item ，要么返回DropItem
    def process_item(self, item, spider):
        if item['link']:
            if item['link'] in self.link_name:
                raise DropItem('重复数据')
            self.link_name.append(item['link'])
            # 存储
            content = json.dumps(dict(item),ensure_ascii=False) + ';\n'
            self.f.write(content)
            return item
        else:
            return DropItem('无效链接')

    def close_spider(self,spider):
        self.f.close()

# 存储在mongo数据库
class MongodbPipeline(object):

    def __init__(self,mongo_url,mongo_db):
        self.mongo_url = mongo_url
        self.mongo_db = mongo_db

    @classmethod
    def from_craler(cls,crawler):
        return cls(
            mongo_url=crawler.settings.get('MONGO_URL'),
            mongo_db=crawler.settings.get('MONGO_DB')
        )

    def open_spider(self,spider):
        self.client = pymongo.MongoClient(self.mongo_url)
        self.db = self.client[self.mongo_db]

    def process_item(self,item,spider):
        self.db['jingdong'].update({'title':item['title']},{'$set':item},True)
        return item

    def close_spider(self,spider):
        self.client.close()