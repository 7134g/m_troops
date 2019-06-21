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


def page_num(url):
    #获取帖子页数
    page = url_open(url)
    num = re.findall('<span class="red">(\d+)',page)

    return int(num[0])
    

def get_img(url):
    #打开贴吧每一页
    for i in range(1,page_num(url)):
        print(i)
        urls = url + '?pn={}'.format(i)
        html = url_open(urls)
        p = r'<img class="BDE_Image" src="([^"]+\.jpg)"'
        imglist = re.findall(p,html)
    
        #保存当前页全部图片
        temp = 1
        for each in imglist:
            filename = each.split('/')[-1]
            urllib.request.urlretrieve(each,filename)
            print('第{}张图片下载完成'.format(temp))
            temp += 1
        time.sleep(1)
        print('第{}页全部下完'.format(i))

## Administrator
def main():
    tt = time.time()
    os.makedirs(r'C:\Users\xincheng\Desktop\meitu')
    os.chdir(r'C:\Users\xincheng\Desktop\meitu')
    #'https://tieba.baidu.com/p/3048618362'
    url = input('请输入贴吧帖子网址：')
    get_img(url)
    flie.close()
    print('总共用时：{}'.format(time.time()-tt))

if __name__ == '__main__':
    main()
