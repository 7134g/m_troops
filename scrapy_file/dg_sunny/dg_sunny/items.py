# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://doc.scrapy.org/en/latest/topics/items.html

from scrapy import Item,Field

from scrapy.loader import ItemLoader
from scrapy.loader.processors import TakeFirst, MapCompose, Join
import re


class DgSunnyItem(Item):
    # define the fields for your item here like:
    # name = scrapy.Field()

    # page
    # LinkExtractor(allow='reply\?page=\d+'),follow = True
    # detali
    # LinkExtractor(allow='http://.*\d+\.shtml'),callback = "detali_parse"

    url = Field()
    # response.url

    title = Field()
    # response.css('.wzy1 .niae2_top::text').extract()[0]

    num = Field()
    # response.css('.wzy1 span[style]::text').re('\d+')[0]

    a_content = Field()
    # response.css('.wzy1 .txt16_3::text').extract()[0]

    g_content = Field()
    # response.css('.wzy1 .txt16_3::text').extract()[2:-1]

class ProgrameItemLoader(ItemLoader):
    default_output_processor = TakeFirst()

def cleanblank_data(value):
    return re.sub(r'\s', '', value)

def get_figure(value):
    return re.findall(r'\d+',value)

def deal_g_content(value):
    return value[2: -1]

class ProgrameItem(Item):


    url = Field()


    title = Field(
        input_processor=MapCompose(cleanblank_data)
                  )


    num = Field(
        input_process=MapCompose(get_figure)
                )


    a_content = Field(
        input_process=MapCompose(cleanblank_data)
                      )


    g_content = Field(
        input_process=MapCompose(cleanblank_data),
        output_process=MapCompose(deal_g_content)
                      )