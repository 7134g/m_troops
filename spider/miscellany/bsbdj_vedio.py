#爬取百思不得姐的视频
import requests
from bs4 import BeautifulSoup
import os
import asyncio
import aiohttp
import aiofiles


#解析网页
async def open_url(url):
    #添加头部信息反爬取
    header = {"User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64)\
    AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 UBrowser/6.2.3964.2 Safari/537.36"}

    connector = aiohttp.TCPConnector(verify_ssl=False)
    async with aiohttp.ClientSession(connector=connector,headers=header) as session:
        async with session.get(url=url) as response:
            return response



#获取视频网页列表
async def get_url(url):
    page_url = []
    response = await open_url(url)
    soup = BeautifulSoup(await response.text(),"html.parser")
    for href in soup.select(".j-r-list-c .j-r-list-c-desc a"):
        page = href['href']
        print(page)
        #print(page)
        href_url = "".join(["http://www.budejie.com" , page])
        page_url.append(href_url)
    #print(page_url)
    return page_url
        
#获取视频地址并将视频下载到文件中
async def down_video(url):
    os.mkdir('mp4')
    os.chdir('mp4')
    page_url = await get_url(url)
    for addres in page_url:
        response = await open_url(addres)
        soup = BeautifulSoup(await response.text(),"html.parser")
        #print(soup)
        #视频地址
        mp4 = soup.select('.j-r-list-c .j-video-c .j-video')[0]['data-mp4']
        #print(mp4)
        file = "".join(["../download/dldl/",mp4.split('/')[-1]])
        async with aiofiles.open(file,"wb") as f:
            #图片和视频用.content,不用.text
            video = await open_url(mp4)
            f.write(video)
              
if __name__ == "__main__":
    count = 10
    loop = asyncio.get_event_loop()
    urls = ["http://www.budejie.com/video/{}".format(index) for index in range(1,count)]
    future = [asyncio.ensure_future(down_video(url)) for url in urls]
    loop.run_until_complete(asyncio.wait(future))


    
