import asyncio
import time
from threading import Thread

async def callback2(timer ,loop):
    await asyncio.sleep(timer/1000)
    print("sleep , {}, {}".format(timer,time.time()))
    return timer

async def callback(loop):
    print("sleep ".format())
    a = []
    loop1 = asyncio.new_event_loop()
    asyncio.set_event_loop(loop1)
    for i in range(10):
        a.append(asyncio.ensure_future(callback2(i * 10, loop)))
    loop1.run_until_complete(asyncio.wait(a))
    # loop1.stop()
    # loop.stop()
    print(loop1.is_running())
    for x in a:
        print(x.result())
    asyncio.set_event_loop(loop)
    loop.stop()
    # if not loop.is_running():
    #     loop.stop()



def main(creat_loop):
    asyncio.set_event_loop(creat_loop)
    creat_loop.run_forever()


tasks = []
creat_loop = asyncio.new_event_loop()
t = Thread(target=main, args=(creat_loop,))
t.start()
asyncio.run_coroutine_threadsafe(callback(creat_loop), creat_loop)
# t.join()
print(tasks)
print(111111)
