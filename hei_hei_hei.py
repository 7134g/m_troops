import urllib.request
import re
import os

pattern = re.compile(r'(?<=<span class="plid">#pl )https://yande.re/post/show/\d+(?=</span>)')
pattern_page = re.compile(r'(?<=src\=")https://files.yande.re/\w+/\w+/yande\.re.+?\.\w+(?=")')

opener = urllib.request.build_opener(urllib.request.ProxyHandler({'代理类型': '代理IP:端口号'})) # 需自行更改, 不用代理可以直接改成urllib.request.ProxyHandler({})
opener.addheaders = [('User-Agent', 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36')]

if not os.path.exists('Picture'):
    os.mkdir('Picture')

os.chdir('Picture')

page = 1
counter = 1
while True:
    for i in pattern.findall(opener.open('https://yande.re/post?page=%d&tags=rating%%3Ae' % (page)).read().decode('utf-8')):
        picture = open('%d.jpg' % (counter), 'wb')
        print('正在下载第%d张图片...' % (counter))
        try:
            picture.write(opener.open(pattern_page.findall(opener.open(i).read().decode('utf-8'))[0]).read())
        except:
            print('第%d张图片下载失败!' % (counter))
        picture.close()
        counter += 1
    page += 1
