import logging
import time
from collections import defaultdict

from pymongo import MongoClient, UpdateOne, DeleteOne
from contextlib import contextmanager


@contextmanager
def close(conn):
    try:
        yield conn
    finally:
        conn.bulk_finally()

logger = logging.getLogger(__name__)


class BaseMongoHandler:
    wcount = 0
    rcount = 0

    def __init__(self):
        client = MongoClient("")
        self.db = client[""]
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
                self.rcount += 1
                if self.rcount % 100000 == 0:
                    logger.info(f"此时读取数为：{self.rcount}")
                yield item

        logger.info("%s集合遍历结束" % collection_name)

    def set_data(self, collection_name, query_builder=None, update_builder=None, upsert=True, is_finish=False):
        """
        传入集合名称，查询条件，更新数据。
        """

        if not is_finish:
            self.bulk[collection_name].append(UpdateOne(query_builder, update_builder, upsert=upsert))
        if len(self.bulk[collection_name]) >= self.MONGOBULK or is_finish:
            if not self.bulk[collection_name]:  # 是否还有剩余数据未提交
                return
            cl = self.db[collection_name]
            s = time.time()
            ret = cl.bulk_write(self.bulk[collection_name])
            e = time.time()
            self.bulk[collection_name] = []
            self.wcount += self.MONGOBULK
            logger.info(f"***{e - s}***, 此时一共写入{self.wcount}个")

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

    def delete(self, col_name, query_builder=None, is_finish=False):
        """
        :param col_name: 表名
        :param data: 目标数据
        :param count: 计数
        """
        col = self.db[col_name]
        if not is_finish:
            self.bulk[col_name].append(DeleteOne(query_builder))

        if len(self.bulk[col_name]) >= self.MONGOBULK or is_finish:
            s = time.time()
            col.bulk_write(self.bulk[col_name])
            e = time.time()
            self.bulk[col_name] = []
            self.wcount += self.MONGOBULK
            logger.info("***%s***, 清除%s个, 当前已操作 %s 个" % (e - s, self.MONGOBULK, self.wcount))


    def bulk_finally(self):
        for col_name in self.bulk:
            cl = self.db[col_name]
            s = time.time()
            cl.bulk_write(self.bulk[col_name])
            e = time.time()
            self.bulk[col_name] = []
            self.wcount += self.MONGOBULK
            logger.info(f"***{e - s}***, 最后一共写入{self.wcount}个")
        logger.info("提交所有剩余数据完毕")
