# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html

from scrapy.pipelines.images import ImagesPipeline
from Douyu.settings import IMAGES_STORE
import scrapy
import os

class DouyuPipeline(ImagesPipeline):

    def get_media_requests(self,item,info):
        # 保存图片
        image_link = item['vertical_src']
        yield scrapy.Request(image_link)

    def item_completed(self,results,item,info):
        # 取出图片路劲名
        image_path = [x['path'] for ok,x in results if ok][0]
        os.rename(IMAGES_STORE + image_path,IMAGES_STORE + item['nickname'] + '.jpg')

        return item
