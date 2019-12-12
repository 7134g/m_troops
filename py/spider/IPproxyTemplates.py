# from db.redisdb import RedisClient
# import config
import redis
import re
import requests
import json
import time
import random


PROXY_POOL_URL = "api地址"
REDIS_HOST = "127.0.0.1"
REDIS_PORT = "6379"
# ip池名
REDIS_PROXY = "proxies"
# ip池容量
REDIS_PROXY_LEN = 5
# ip失效时间
PROXY_TIMEOUT = 300

class RedisClient:
    def __init__(self, db=1):
        self.db = redis.StrictRedis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True, db=db)

    def add(self, proxy: str, score: int):
        """
        添加代理，设置分数为最高
        :param proxy: 代理
        :param score: 分数
        :return: 添加结果
        """
        if not re.match('\d+\.\d+\.\d+\.\d+\:\d+', proxy):
            print('代理不符合规范', proxy, '丢弃')
            return
        if not self.db.zscore(REDIS_PROXY, proxy):
            return self.db.zadd(REDIS_PROXY, {proxy: score})

    def redis_get(self, key: str, index=-1):
        """
        默认获取score最小的任务
        :return: str
        """
        result = self.db.zrevrange(key, 0, 100000)
        if len(result):
            return result[index]
        else:
            raise Exception("没有任务可以获取")

class ProxyClient:
    def __init__(self):
        self.proxy_redis = RedisClient(db=2)

    def get_ip(self):
        return self.proxy_redis.redis_get(key=REDIS_PROXY, index=random.randint(0, REDIS_PROXY_LEN-1))

    def pull_ip(self):
        # 开始获取新的一批ip
        proxy_data = requests.get(PROXY_POOL_URL)
        json_data = json.loads(proxy_data.text)
        now_time = int(time.time() * 1000)
        # print(type(json_data))
        print(json_data)
        for proxy in json_data:
            ip_content = ":".join([str(proxy["http_ip"]), str(proxy["http_port"])])
            self.proxy_redis.add(ip_content, now_time)


def check_timeout(proxy):

    redis_ = proxy.proxy_redis
    timeout_task = []  # 超时任务
    taskList = redis_.redis_Rall(REDIS_PROXY, scoreTpye=True)  # 读取当前键所有[(value,score)]

    for taskData in taskList:
        taskStr = taskData[0]
        parent = taskData[1]
        if int(time.time() * 1000) - parent > PROXY_TIMEOUT:
            timeout_task.append(taskStr)

    # 删除redis表中的对应的值
    redis_.redis_Dtimeout(REDIS_PROXY, timeout_task)
    while (redis_.redis_count(REDIS_PROXY) < REDIS_PROXY_LEN):
        proxy.pull_ip()
        print("添加ip")


if __name__ == '__main__':
    x = ProxyClient()
    x.pull_ip()
