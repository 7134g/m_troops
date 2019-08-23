import re
import urllib.request
import os
import time


def url_open(url):
    #打开连接
    req = urllib.request.Request(url)
    req.add_header("User-Agent","Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
    page = urllib.request.urlopen(req)
    html = page.read().decode('utf-8')

    return html

def cbk(a,b,c):
    per = 100.0 * a * b / c
    if per > 100:
        per = 100
    print('%.2f%%' % per)
    print('下载完成')

def page_num(url):
    #获取帖子页数
    page = url_open(url)
    num = re.findall('<span class="red">(\d+)',page)

    return int(num[0])
    

def get_img(url):
    #打开贴吧每一页
    for i in range(1,page_num(url)):
        urls = url + '?pn={}'.format(i)
        html = url_open(url)
        p = r'<img class="BDE_Image" src="([^"]+\.jpg)"'
        imglist = re.findall(p,html)
    
        #保存当前页全部图片
        for each in imglist:
            filename = each.split('/')[-1]
            urllib.request.urlretrieve(each,filename,cbk)
        time.sleep(1)

def main():
    path = os.getcwd()+'meitu'
    os.makedirs(path)
    os.chdir(path)
    #'https://tieba.baidu.com/p/3048618362'
    url = input('请输入贴吧帖子网址：')
    get_img(url)


if __name__ == '__main__':
    main()
