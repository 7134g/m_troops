import time
import os
from urllib import request
from lxml import etree

def get_html(url):
    headers = {'User-Agent': r'Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:63.0) Gecko/20100101 Firefox/63.0',
               'Connection': r'keep-alive',
               'Referer': r'http://www.itcast.cn/channel/teacher.shtml'}
    req = request.Request(url, headers=headers)
    page = request.urlopen(req).read().decode('utf-8')
    html = etree.HTML(page)
    return html


def download(url_3,m):
        html_3 = get_html(url_3)
        # 得到小说名称
        title_novel = html_3.xpath('//div/div/span/a/text()')[1]
        # print(title_novel)
        #得到小说每一个章节的名称
        chapter = html_3.xpath('//div/h1/text()')[0]
        # print(chapter)
        #得到每一个章节小说的内容
        content = html_3.xpath('//body/div/div/p/text()')
        if m == 'o':
            folder = r'旧版/%s/' % title_novel
            if (not os.path.exists(folder)):
                os.makedirs(folder)
        elif m == 'n':
            folder = '新修版/%s/' % title_novel
            if (not os.path.exists(folder)):
                os.makedirs(folder)
        else:
            folder = '修订版/%s/' % title_novel
            if (not os.path.exists(folder)):
                os.makedirs(folder)

        filename = folder + chapter + '.txt'
        with open(filename, 'a', encoding="utf-8") as f:
            f.write(chapter + '\n')
        for j in content:
            with open(filename, 'a', encoding="utf-8") as f:
                f.write(j + '\n')
        print('正在下载，请稍后....')


def main(url):
    url_2 = 'http://www.jinyongwang.com' + url
    html_2 = get_html(url_2)
    data_urls = html_2.xpath('//ul[@class="mlist"]/li/a/@href')
    for i in data_urls:
        url_3 = 'http://www.jinyongwang.com' + i
        m = i[1]
        download(url_3,m)


if __name__ == '__main__':
    #需要获取的第一个网址
    start_time = time.time()
    star_url = 'http://www.jinyongwang.com/book/'
    html_1 = get_html(star_url)
    urls = html_1.xpath('//ul[@class="list"]/li/p[@class="title"]/a/@href')
    main(urls)
    end_time = time.time()
    print(end_time-start_time)
