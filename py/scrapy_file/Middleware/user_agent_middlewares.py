# -*- coding: utf-8 -*-
from fake_useragent import UserAgent
from scrapy.http import Headers


class RandomUserAgentMiddlware(object):
    #随机跟换user-agent
    def __init__(self,crawler):
        super(RandomUserAgentMiddlware,self).__init__()
        self.ua = UserAgent(use_cache_server=False)
        self.ua_type = crawler.settings.get('RANDOM_UA_TYPE','random')#从setting文件中读取RANDOM_UA_TYPE值

    @classmethod
    def from_crawler(cls,crawler):
        return cls(crawler)

    def process_request(self, request, spider):
        def get_ua():
            return getattr(self.ua,self.ua_type)
        headers = {
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
            'Accept-Language': 'zh-CN,zh;q=0.9',
            'Connection': 'keep - alive',
            'referer': 'https://search.jd.com/Search?keyword=%E5%9B%BE%E4%B9%A6&enc=utf-8&wq=%E5%9B%BE%E4%B9%A6&page=1'
        }
        headers['User_Agent'] = get_ua()
        request.headers = Headers(headers)
        print(request.headers)
