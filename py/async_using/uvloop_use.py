import uvloop
import asyncio

uvloop.install()
new_loop = uvloop.new_event_loop()
# 注册进当前进程
asyncio.set_event_loop(loop)

