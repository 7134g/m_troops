import datetime
import logging
import time
from collections import defaultdict
from pymongo import MongoClient, UpdateOne, ReplaceOne, DeleteOne

from py.sql_and_nosql.tools import get_project_settings

logger = logging.getLogger(__name__)


class BaseMongoHandler:
    def __init__(self, db_name):
        setting = get_project_settings()
        mongo_uri = setting.MONGO_URI
        client = MongoClient(mongo_uri)
        self.db = client[db_name]
        # self.db.collection_names()
        self.bulk = defaultdict(list)
        # 批量操作数量
        self.MONGOBULK = 100

    def get_db(self):
        return self.db

    def get_cursor(self, collection_name, query):
        """
        传入集合名称和查询条件，遍历集合
        """
        col = self.db[collection_name]
        with col.find(query, no_cursor_timeout=True).batch_size(1000) as cursor:
            for item in cursor:
                yield item

        logger.info("%s集合遍历结束" % collection_name)

    def set_data(self, collection_name, query_builder=None, update_builder=None, is_finish=False):
        """
        传入集合名称，查询条件，更新数据。
        """

        if not is_finish:
            self.bulk[collection_name].append(UpdateOne(query_builder, update_builder, upsert=True))
        if len(self.bulk[collection_name]) >= self.MONGOBULK or is_finish:
            if not self.bulk[collection_name]:  # 是否还有剩余数据未提交
                return
            cl = self.db[collection_name]
            s = time.time()
            cl.bulk_write(self.bulk[collection_name])
            e = time.time()
            self.bulk[collection_name] = []
            logger.debug("***********%s***********, 入库%s个" % (e - s, self.MONGOBULK))

    def get_find_one(self, collection_name, query_builder):
        """
        通过名字查询出公司内容，返回整条数据
        """
        s = time.time()
        ret = self.db[collection_name].find_one(query_builder)
        e = time.time()
        logger.info("读取mongo时间为：%s %s" % (e - s, query_builder))
        if ret:
            return ret

        return None

    def replace(self, col_name, query_builder, data, count, is_finish=False):
        """
        :param col_name: 表名
        :param data: 目标数据
        :param count: 计数
        """
        col = self.db[col_name]
        cur_time = datetime.datetime.utcnow()
        del data['_id']
        data['updateTime'] = cur_time

        if not is_finish:
            self.bulk.append(ReplaceOne(query_builder, data, upsert=True))

        if len(self.bulk[col_name]) >= self.MONGOBULK or is_finish:
            s = time.time()
            col.bulk_write(self.bulk)
            e = time.time()
            self.bulk = []
            print("***%s***, 替换%s个, 当前已操作 %s 个" % (e - s, self.MONGOBULK, count))

    def delete(self, col_name, query_builder, count, is_finish=False):
        col = self.db[col_name]
        if not is_finish:
            self.bulk[col_name].append(DeleteOne(query_builder))

        if len(self.bulk[col_name]) >= self.MONGOBULK or is_finish:
            s = time.time()
            col.bulk_write(self.bulk[col_name])
            e = time.time()
            self.bulk = []
            print("***%s***, 清除%s个, 当前已操作 %s 个" % (e - s, self.MONGOBULK, count))
