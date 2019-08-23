# -*- coding: utf-8 -*-
from scrapy import Spider,Request
import json
from zhihu.items import *


class FensSpider(Spider):
    name = 'fens'
    allowed_domains = ['www.zhihu.com']
    start_urls = ['http://www.zhihu.com/']

    # 初始人
    start_user = 'zhang-jia-wei'
    user_url = 'https://www.zhihu.com/api/v4/members/{user}?include={include}'
    user_query = 'allow_message%2Cis_followed%2Cis_following%2Cis_org%2Cis_blocking%2Cemployments%2Canswer_count%2Cfollower_count%2Carticles_count%2Cgender%2Cbadge%5B%3F(type%3Dbest_answerer)%5D.topics'

    # 粉丝列表
    fans_url = 'https://www.zhihu.com/api/v4/members/{user}/followers?include={include}&offset={offset}&limit={limit}'
    fans_query = 'data%5B*%5D.answer_count%2Carticles_count%2Cgender%2Cfollower_count%2Cis_followed%2Cis_following%2Cbadge%5B%3F(type%3Dbest_answerer)%5D.topics'

    # 关注列表
    follows_url = 'https://www.zhihu.com/api/v4/members/{user}/followees?include={include}&offset={offset}&limit={limit}'
    follows_query = 'data%5B*%5D.answer_count%2Carticles_count%2Cgender%2Cfollower_count%2Cis_followed%2Cis_following%2Cbadge%5B%3F(type%3Dbest_answerer)%5D.topics'

    def start_requests(self):
        # 用户页
        yield Request(self.user_url.format(user=self.start_user,include = self.user_query),self.parse_user)
        # 关注页
        yield Request(self.follows_url.format(user=self.start_user,include=self.follows_query,offset=0,limit=20),callback = self.parse_follows)
        # 粉丝页
        yield Request(self.fans_url.format(user=self.start_user,include=self.fans_query,offset=0,limit=20),callback = self.parse_fans)


    def parse_user(self, response):
        # 解析user详细信息
        result = json.loads(response.text)
        item = UserItem()
        for field in item.fields:
            if field in result.keys():
                item[field] = result.get(field)
        yield item
        # 获取关注列表
        yield Request(self.follows_url.format(user=result.get('url_token'),include=self.follows_query,limit=20,offset=0),callback=self.parse_follows)
        # 获取粉丝列表
        yield Request(self.fans_url.format(user=self.start_user, include=self.fans_query, offset=0, limit=20),callback=self.parse_fans)

    def parse_follows(self,response):
        results = json.loads(response.text)

        if 'data' in results.keys():
            for result in results.get('data'):
                yield Request(self.user_url.format(user=result.get('url_token'),include=self.user_query),callback=self.parse_user)

        if 'paging' in results.keys() and results.get('paging').get('is_end')==False:
            next_page = results.get('paging').get('next')
            yield  Request(url = next_page,callback= self.parse_follows)



    def parse_fans(self,response):
        results = json.loads(response.text)

        if 'data' in results.keys():
            for result in results.get('data'):
                yield Request(self.user_url.format(user=result.get('url_token'),include=self.user_query),callback=self.parse_user)

        if 'paging' in results.keys() and results.get('paging').get('is_end')==False:
            next_page = results.get('paging').get('next')
            yield  Request(url = next_page,callback= self.parse_fans)

