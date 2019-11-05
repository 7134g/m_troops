import asyncio
import time
from threading import Thread
# import uvloop

now = lambda :time.time()


def start_loop(loop):
    asyncio.set_event_loop(loop)
    loop.run_forever()

async def do_some_work(x):
    print('Waiting {}'.format(x))
    await asyncio.sleep(x)
    print('Done after {}s'.format(x))


start = now()
# new_loop = uvloop.new_event_loop()
new_loop = asyncio.new_event_loop()
t = Thread(target=start_loop, args=(new_loop,))
t.start()
print('TIME: {}'.format(time.time() - start))

for i in range(10):#协程开启十万个
    asyncio.run_coroutine_threadsafe(do_some_work(i*10), new_loop)
t.join() # 阻塞主线程
for i in range(10,20):#协程开启十万个
    asyncio.run_coroutine_threadsafe(do_some_work(i*10), new_loop)