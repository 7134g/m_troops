# coding: utf-8
import requests
from random import choice
# import importlib
import traceback
import time
import json

from Com import baidu_id_secret
# importlib.reload(baidu_id_secret)


class BaiduNLP(object):
    def __init__(self):
        pass

    def get_id_secret(self):
        # todo 获取 client_id 和 client_secret
        id_secret = choice(baidu_id_secret.id_secret_list)
        client_id = id_secret['client_id']
        client_secret = id_secret['client_secret']
        return client_id, client_secret

    def get_access_token(self, client_id, client_secret):
        access_token_url = 'https://aip.baidubce.com/oauth/2.0/token'
        access_token_params = {
            'grant_type': 'client_credentials',
            'client_id': client_id,
            'client_secret': client_secret,
        }
        access_token_headers = {'Content-Type': 'application/json; charset=UTF-8'}
        access_token_resp = requests.get(access_token_url, params=access_token_params, headers=access_token_headers, timeout=5)
        access_token_json = access_token_resp.json()
        # print(access_token_json)
        return access_token_json

    def get_topic(self, access_token, title, content):
        if len(title) > 30:
            title = title[0:30]
        if len(content) > 5000:
            content = content[0:5000]
        # url = 'https://aip.baidubce.com/rpc/2.0/nlp/v1/topic?access_token=' + access_token
        url = 'https://aip.baidubce.com/rpc/2.0/nlp/v1/topic'
        tag_params = {
            'access_token': access_token
        }
        post_data = {
            "title": title,
            "content": content
        }
        headers = {
            'Content-Type': 'application/json'
        }

        resp = requests.post(url, headers=headers, params=tag_params, json=post_data,verify=False, timeout=5)
        # resp = requests.post(url, headers=headers, params=params, data=post_data)
        resp.encoding = 'gbk'
        topic_json = resp.json()
        return topic_json

    def get_keyword(self, access_token, title, content):
        # url = 'https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=' + access_token
        url = 'https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword'
        tag_params = {
            'access_token': access_token
        }
        post_data = {
            "title": title,
            "content": content
        }
        headers = {
            'Content-Type': 'application/json'
        }

        resp = requests.post(url, headers=headers, params=tag_params, json=post_data, timeout=5,verify=False)
        # resp = requests.post(url, headers=headers, params=params, data=post_data)
        resp.encoding = 'gbk'
        keyword_json = resp.json()
        return keyword_json


# todo access_token 的有效期为30天，25天更新一次
def generate_access_token_list():
    access_token_list = []
    bf = BaiduNLP()
    for id_secret in baidu_id_secret.id_secret_list:
        client_id = id_secret['client_id']
        client_secret = id_secret['client_secret']

        retry = 3
        while retry > 0:
            try:
                access_token_json = bf.get_access_token(client_id, client_secret)
                access_token = access_token_json.get('access_token', '')
                if not access_token:
                    raise Exception(client_id + ', ' + client_secret+', ' + 'get_access_token error: ' + str(access_token_json))
                access_token_list.append(access_token)
                break
            except:
                retry -= 1
                time.sleep(0.1)
                if retry == 0:
                    print(traceback.format_exc())
        # time.sleep(0.1)

    print(access_token_list)
    print(len(access_token_list))
    with open('access_tokens.py', 'w', encoding='utf-8') as fd:
        fd.write('access_token_list = ' + json.dumps(access_token_list, ensure_ascii=False))


if __name__ == '__main__':
    # generate_access_token_list()
    pass


