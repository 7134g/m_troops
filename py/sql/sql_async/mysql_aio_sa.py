import asyncio

import sqlalchemy as sa

from aiomysql.sa import create_engine




async def go(loop):
    engine = await create_engine(user='root', db='test_pymysql',
                                 host='127.0.0.1', password='fxj123', loop=loop
                                 # , autocommit=True
                                 )
    async with engine.acquire() as conn:
        await conn.execute(user.insert().values(val=40))
        await conn.execute(user.insert().values(val=50))

        row = await conn.execute(user.select().where(user.columns.val > 20))
        a = await row.fetchall()
        # async for i in row.fetchall():
        #     print(i)
        async for i in a:
            print(i)
        # print(type(a))
        # print(a[0].val)
        # print(dict(a[0]))

    engine.close()
    await engine.wait_closed()


if __name__ == '__main__':
    metadata = sa.MetaData()
    user = sa.Table('user', metadata,
                   sa.Column('id', sa.Integer, primary_key=True),
                   sa.Column('val', sa.String(255)))
    loop = asyncio.get_event_loop()
    loop.run_until_complete(go(loop))
