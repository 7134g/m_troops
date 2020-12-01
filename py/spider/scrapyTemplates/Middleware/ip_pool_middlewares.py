# -*- coding:utf8 -*-

import requests
import base64


class Get_ip(object):
    def __init__(self):
        self.PROXY_POOL_URL = 'http://localhost:5555/random'

    def get_proxy(self):
        try:
            response = requests.get(self.PROXY_POOL_URL)
            if response.status_code == 200:
                return response.text
        except ConnectionError:
            return None


class Ip_Pool_Middlewares(object):

    def __init__(self):
        self.PROXIES = Get_ip().get_proxy()


    # 传进带密码的
    def process_request(self,request,spider):

        if self.PROXIES is dict and self.PROXIES['user_password'] != None:
            base64_userpasswd = base64.b64encode(self.PROXIES['user_password'])
            request.meta['proxy'] = "http://" + self.PROXIES['proxy']
            request.headers['Proxy-Authorization'] = 'Basic' + base64_userpasswd
        else:
            request.meta['proxy'] = "http://"+ self.PROXIES