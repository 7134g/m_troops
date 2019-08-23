# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html

import json
import os
import pymongo

class ZhihuPipeline(object):

    def __init__(self):
        self.file = open(os.getcwd()+'\data\jingdong.json','w')


    def process_item(self, item, spider):
        content = json.dumps(dict(item),ensure_ascii=False)+';\n'
        self.file.write(content)
        return item

    def close_spider(self,spider):
        self.file.close()

class MongoPipeline(object):

    def __init__(self,mongo_uri,mongo_db):
        self.mongo_uri = mongo_uri
        self.mongo_db = mongo_db

    @classmethod
    def from_crawler(cls,crawler):
        return cls(
            mongo_uri=crawler.settings.get('MONGO_URI'),
            mongo_db = crawler.settings.get('MONGO_DB')
        )
    def open_spider(self,spider):
        self.client = pymongo.MongoClient(self.mongo_uri)
        self.db = self.client[self.mongo_db]

    def close_spider(self,spider):
        self.client.close()

    def process_item(self,item,spider):
        self.db['user'].update({'url_token':item['url_token']},{'$set':item},True)
        return item

