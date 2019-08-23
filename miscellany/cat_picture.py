# import urllib.request
# response = urllib.request.urlopen('http://placekitten.com/g/500/600')
# cat_img = response.read()
#
# with open('cat_500_600.jpg','wb') as f:
#     f.write(cat_img)
#

import asyncio
import aiohttp
import aiofiles



async def get_img(w, h, n):
    url_img = '{}/{}/{}'.format(url_base,w, h)
    connector = aiohttp.TCPConnector()
    async with aiohttp.ClientSession(connector=connector) as session:
        async with session.get(url_img) as response:
            async with aiofiles.open('{}_{}_{}.jpg'.format(w, h, n),'wb') as f:
                f.write(await response.text())


if __name__ == '__main__':
    print("输入图片宽：")
    w = input()
    print("输入图片高：")
    h = input()
    print("输入0彩色，1黑白")
    ch = input()

    if ch == "0" or ch == 0:
        url_base = "http://placekitten.com"
    else:
        url_base = "http://placekitten.com/g"

    loop = asyncio.get_event_loop()
    future = [asyncio.ensure_future(get_img(w,h,n)) for n in range(10)]
    loop.run_until_complete(asyncio.wait(future))