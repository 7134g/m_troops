import asyncio
import aiomysql


class MysqlOptAsync(object):
    def __init__(self, loop):
        self.host = "127.0.0.1"
        self.port = 3306
        self.user = "root"
        self.password = "fxj123"
        self.db = "test_pymysql"

        self._pool = None  # 存放mysql连接池
        self._loop = loop  # mysql连接池的事件循环基础

    async def pool(self):
        if not self._pool:
            print('开始生成链接池')
            self._pool = await aiomysql.create_pool(host=self.host, port=self.port, user=self.user,
                                                    password=self.password, db=self.db, loop=self._loop)
            print(self._pool)
        return self._pool

    async def insertOpt(self, data=None):
        # 获取连接池对象
        async with self._pool.acquire() as conn:
            # 建立光标
            async with conn.cursor() as cur:
                sql = 'insert into user(val) value(%s);'
                sql2 = 'SELECT * FROM USER WHERE val=21;'
                try:
                    # 生成需要执行的sql语句
                    # await cur.execute(sql, data)
                    a = await cur.execute(sql2)
                    print(a)
                    # 提交操作
                    await conn.commit()
                except Exception as e:
                    print('cuo')
                    print(e)
                    await conn.rollback()
                # await cur.commit()
                # print(await cur.last_id())


async def insert_data(obj_str=None):
    for i in range(21, 40):
        await obj_str.insertOpt(data=(i,))


async def main(loop):
    mysql = MysqlOptAsync(loop=loop)
    await mysql.pool()
    await insert_data(obj_str=mysql)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main(loop))
