# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://doc.scrapy.org/en/latest/topics/items.html

from scrapy import Item,Field


class UserItem(Item):
    # define the fields for your item here like:
    # name = scrapy.Field()

    answer_count = Field()
    articles_count = Field()

    avatar_url = Field()
    avatar_url_template = Field()

    follower_count = Field()
    gender = Field()

    url_token = Field()
    headline = Field()


    id = Field()
    name = Field()
    type = Field()
    url = Field()
    user_type = Field()


