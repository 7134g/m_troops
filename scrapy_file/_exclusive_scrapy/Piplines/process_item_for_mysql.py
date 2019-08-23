# ********************************************* 同步机制 ******************************************************
# piplines.py

import MySQLdb

class MySQLPipline(object):
    def __init__(self):
        self.conn = MySQLdb.connect(host = '127.0.0.1',user = 'root',password = 'mysql',dbname = 'jobdb',charset = 'utf8',use_unicode=True)
        self.cursor = self.conn.cursor()

    def process_item(self,item,spider):
        # job_article是数据表名
        insert_sql = '''
            insert into job_article(title,url,data,nums)
            VALUES (%s, %s, %s, %s)
        '''
        self.cursor.execute(insert_sql , (item['title'],item['url'],item['data'],item['nums']))
        self.conn.commit()

# setting.py

ITEM_PIPLINE = {
       'jingdong.pipelines.MySQLPipeline':1
}


# ********************************************* 异步机制 ******************************************************
# piplines.py

import MySQLdb
import MySQLdb.cursors
from twisted.enterprise import adbapi

class MySQLTwistedPipline(object):

    def __init__(self,dbpool):
        self.dbpool = dbpool

    @classmethon
    def from_setting(cls,setting):
        dbparms = dict(
            host = setting["MYSQL_HOST"],
            db = setting["MYSQL_DBNAME"],
            user = setting["MYSQL_USER"],
            passwd = setting["MYSQL_PASSWORD"],
            charset = 'utf8',
            cursorclass = MySQLdb.cursors.DictCursor,
            use_unicode = True
            )
        # args普通参数,*args可变参数(变元组),**kwargs可变参数(变字典)
        self.dbpool = adbapi.ConnectionPool("MySQLdb",**dbparms)
    return cls(dbpool)

    def process_item(self,item,spider):
        # 使用twisted将mysql插入变成异步执行
        query = self.dbpool.runInteraction(self.do_insert,item)
        query.addErrback(self.handle_error,item,spider)   # 处理异常

    def do_insert(self,cursor, item):
        # job_article是数据表名
        insert_sql = '''
            insert into job_article(title,url,data,nums)
            VALUES (%s, %s, %s, %s)
        '''
        cursor.execute(insert_sql,(item['title'],item['url'],item['data'],item['nums']))

    # 处理异步的异常的函数
    def handle_error(self, failure, item, spider):
        print(failure)


# setting.py

MYSQL_HOST = 'localhost'
MYSQL_DBNAME = 'jobdb'
MYSQL_USER = 'root'
MYSQL_PASSWORD = 'mysql'