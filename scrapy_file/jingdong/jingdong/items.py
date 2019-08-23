# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# http://doc.scrapy.org/en/latest/topics/items.html

import scrapy


class JingdongItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()

    # 商品ID
    Id = scrapy.Field()
    catId = scrapy.Field()
    #商品链接
    link = scrapy.Field()
    #标题
    title = scrapy.Field()
    #作者
    writer = scrapy.Field()
    #原价
    o_price = scrapy.Field()
    #实际价格
    n_price = scrapy.Field()
    #电子书价格
    e_price = scrapy.Field()
    #评论数
    comment = scrapy.Field()

