#爬取百思不得姐的视频
import requests
from bs4 import BeautifulSoup
import os

#解析网页
def open_url(url):
    #添加头部信息反爬取
    header = {"User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64)\
    AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 UBrowser/6.2.3964.2 Safari/537.36"}

    res = requests.get(url,headers=header)
    
    #print(res.text)
    return res



#获取视频网页列表
def get_url(url):
    page_url = []
    soup = BeautifulSoup(open_url(url).text,"html.parser")
    for href in soup.select(".j-r-list-c .j-r-list-c-desc a"):
        page = href['href']
        print(page)
        #print(page)
        href_url = "http://www.budejie.com" + page
        page_url.append(href_url)
    #print(page_url)
    return page_url
        
#获取视频地址并将视频下载到文件中
def down_video(url):
    os.mkdir('mp4')
    os.chdir('mp4')
    count = 0
    page_url = get_url(url)
    for addres in page_url:
        
        soup = BeautifulSoup(open_url(addres).text,"html.parser")
        #print(soup)
        #视频地址
        mp4 = soup.select('.j-r-list-c .j-video-c .j-video')[0]['data-mp4']
        #print(mp4)
        file = mp4.split('/')[-1]
        with open(file,"wb") as f:
            #图片和视频用.content,不用.text
            video = open_url(mp4).content
            f.write(video)
              
if __name__ == "__main__":
    i = 1
   # 设置爬取 9 页，从第一页开始（可以自己设置）
    while i<10:
        url = "http://www.budejie.com/video/" + str(i)
        down_video(url)
        i += 1

    
