# 以子线程方式启动事件循环，主线程用于往其中动态添加任务
# 内存峰值 520000, 花费 20 秒
import asyncio
import datetime , time
from threading import Thread
# import uvloop

t_now = time.time()
now = datetime.datetime.now()

def start_loop(loop):
    asyncio.set_event_loop(loop)
    loop.run_forever()

async def do_some_work(x):
    print('Waiting {}'.format(x))
    await asyncio.sleep(10)
    print('TIME: {} - {} == {}'.format(now, datetime.datetime.now(), time.time() - t_now))


def tasks(count):
    for i in range(count):  # 协程开启十万个
        asyncio.run_coroutine_threadsafe(do_some_work(i), new_loop)
    # new_loop.call_soon_threadsafe(new_loop.stop)
# new_loop = uvloop.new_event_loop()
new_loop = asyncio.new_event_loop()
t = Thread(target=start_loop, args=(new_loop,))
t.start()


tasks(100000)


