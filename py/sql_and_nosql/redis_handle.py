import redis
from py.sql_and_nosql.tools import get_project_settings


class RedisClient(object):
    def __init__(self):
        setting = get_project_settings()
        REDIS_PARAMS = setting.REDIS_PARAMS
        HOST = REDIS_PARAMS['host']
        PORT = REDIS_PARAMS['port']
        PARAMS = REDIS_PARAMS['password']
        REDIS_DB = REDIS_PARAMS['db']
        self.db = redis.StrictRedis(host=HOST,
                                    port=PORT,
                                    password=PARAMS,
                                    decode_responses=True,
                                    db=REDIS_DB)


if __name__ == '__main__':
    r = RedisClient()
    db = r.db