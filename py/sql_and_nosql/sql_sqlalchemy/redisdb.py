import redis
import json
from random import choice
import re
from dateutil.parser import parse
from datetime import datetime

from config import REDIS_HOST, REDIS_PORT, REDIS_PROXY


class PoolEmptyError(Exception):

    def __init__(self):
        Exception.__init__(self)

    def __str__(self):
        return repr('代理池已经枯竭')


class RedisClient(object):
    def __init__(self, db=1):
        """
        初始化
        :param host: Redis 地址
        :param port: Redis 端口
        :param password: Redis密码
        plat.xiecheng.xiecheng_seating
        """
        self.db = redis.StrictRedis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True,db=db)

    def close(self):
        return self.db.connection_pool.disconnect()

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

    def redis_W(self, results: list, set_score=None):
        """
        写入数据
        """
        print("开始存储数据：{}".format(results))
        # core is value
        key = results[0]
        value = json.dumps(results[1], ensure_ascii=False)
        score = results[2]
        if set_score != None:
            score = set_score
            print("redis执行成功，分数为time")

        self.db.zadd(key, {value: score})
        print("redis执行成功，分数为{}".format(score))

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

    def redis_Dtimeout(self, key: str, timeout_task: list):
        """
        删除某一键中的，部分value值
        :param key: 键名
        :param check_time: 超时时间
        :return:
        """
        for value in timeout_task:
            self.redis_Donce(key, value)
        delete_value = len(timeout_task)
        print("{}删除了{}个".format(key, delete_value))
        return delete_value

    def redis_Donce(self, key: str, value: str):
        """
        根据键值对，删除
        :param key: 键名
        :param value: 值
        :return: 成功返回1，失败0
        """
        return self.db.zrem(key, value)

    def redis_Dattr(self, key: str, parent: int):
        """
        根据对比parent和up数值判断是否为重复任务
        若是重复任务则删除旧任务
        :param key : 键名
        :param parent : value值中的parent字段，是时间戳
        :return:
        """
        all_value = self.db.zrange(key, 1, 10**10)
        for value in all_value:
            valueReturnDict = json.loads(value)
            if parent == valueReturnDict["up"]:
                self.redis_Donce(self, key, value)

    def redis_Ddate(self, key: str):
        all_value = self.db.zrange(key, 1, 10 ** 10)
        for value in all_value:
            valueReturnDict = json.dumps(value)
            now = datetime.now()
            target = parse(valueReturnDict["start"])
            if (now-target).seconds > 86400:
                self.redis_Donce(self, key, value)

    def redis_Rall(self, key: str, scoreTpye=False):
        '''
        读取表中全部
        :return []  or [()]
        '''
        return self.db.zrange(key, start=0, end=10**5, withscores=scoreTpye)

    def r_score(self, key: str, value: int):
        """
        获取键值对应的分数
        :param key: 键
        :param value: 值
        :return: int
        """
        return self.db.zscore(key, value)

    def r_score_count(self, key: str, score: int):
        """
        获取相同分数的数量
        :param key: 键
        :param min: 最小
        :param max: 最大
        :return: list
        """
        min_ = max_ = score
        return self.db.zcount(key, min_, max_)

    def redis_count(self, key: str):
        """
        获取数量
        :return: 数量
        """
        return self.db.zcard(key)

    def random_proxy(self):
        """
        随机获取有效代理，首先尝试获取最高分数代理，如果不存在，按照排名获取，否则异常
        :return: 随机代理
        """
        result = self.db.zrangebyscore("proxies", 0, 1000000)
        if len(result):
            return choice(result)
        else:
            result = self.db.zrevrange("proxies", 0, 1000000)
            if len(result):
                return choice(result)
            else:
                raise PoolEmptyError


if __name__ == '__main__':
    r = RedisClient(db=1)
    # num = randint(1,100) + 1465164366000
    # upt = int(time.time()*1000)-2000000
    result = r.r_score("20190806tpeosa3K721",'')
    print(result)
    # for x in result:
    #     print(x[1])
        # print(x[0])
    # print(len(result))
    # print(result)