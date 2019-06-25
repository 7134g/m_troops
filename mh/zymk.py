
#注意，要用到谷歌无头浏览器哦，，可以自己去安装，
#教程https://www.jianshu.com/p/11d519e2d0cb


import os
import re
import requests
from tkinter import *
from lxml import etree

# 导入chrome无头浏览器
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

'''
下面的函数作用及功能：
1，获取到要下载书籍的最新（最大）章节，以便后面遍历章节需要。
2，获取到每一章节的名字与url，
3，遍历每一章节的url，获取到每一章节一共有多少分页（话）
4，调用函数【url_key(url) # 获取到最新章节图片的url，分析规则。】
5，调用函数【save_img(ZhangJie, pn_high, book_names, html) # 进行图片保存】
'''


def all_all(url, book_names):
    headers = {
        'User-Agent': 'Mozilla/5.0(Windows NT 10.0; Win64; x64)AppleWebKit/537.36(KHTML,like Gecko)Chrome/70.0.3538.77Safari/537.36',
    }
    req = requests.get(url, headers=headers)  # .content#proxies=
    html = etree.HTML(req.text)

    pn_list = []  # 获取最新章节，或者是最大的章节
    url_keys = []
    for i in range(1, 5):  # 获取前5个章节
        pn2 = html.xpath('//*[@id="chapterList"]/li[{}]/a/@title'.format(i))
        pn2 = re.findall('\d+', str(pn2))
        if len(pn2) > 0:
            # print(pn2[0],333333333333)
            url_k = html.xpath('//*[@id="chapterList"]/li[{}]/a/@href'.format(i))  # 获取到最大章节的url
            url_keys.append(url_k[0])  # 将最大章节URL 返回
            pn_list.append(pn2[0])  # 将最大章节 返回
    pn = max(pn_list)  # 返回最大的章节

    confirmLabel.delete(0, END)  # 清空文本框
    confirmLabel.insert(END, '{}一共有{}章'.format(book_names, pn), '\n', )  # 打印到GUI界面
    confirmLabel.insert(END, '图片将保存在本程序运行的文件夹，请注意查看哦', '\n')  # 打印到GUI界面
    window.update()  # 刷新文本框显示内容

    htmls = url_key(url + url_keys[pn_list.index(pn)])  # 查找pn在列表里面的第几个，相应的取第几个URL

    save_img(int(pn), url, book_names, htmls, html)


# 获取到最新章节图片的url，分析规则。
def url_key(url):
    chrome_options = Options()
    chrome_options.add_argument("--headless")
    driver = webdriver.Chrome(chrome_options=chrome_options)

    driver.get(url)
    webpage = driver.page_source
    driver.close()
    return webpage  # 返回最新一章的页面内容，用作查看url规则


# 保存图片函数
def save_img(ZhangJie, url, book_names, htmls, html):
    htmlsf = etree.HTML(htmls)
    req = htmlsf.xpath('/html/body/div[1]/div[1]/div[2]/div[1]/img/@src')
    urls = str(req)
    s = re.findall('%E8%AF%9D.*%2F1', urls)
    url_key = re.sub('%E8%AF%9D|%2F1|\[\'|\'|\]', '', str(s))
    url_tou = re.sub('2F\d+\S+|\[\'', '', urls)
    for i in range(1, ZhangJie + 1):
        # for pn in range(1, int(pn_high[i - 1]) + 1):
        # 获取到每一章节的名字与url
        pnx = html.xpath(
            '//*[@id="chapterList"]/li[{}]/a/@title|//*[@id="chapterList"]/li[{}]/a/@href'.format(i, i))
        if str(pnx).find('话') >= 0:
            req = requests.get(url + pnx[0])
            htmlsa = etree.HTML(req.text)
            url_f = htmlsa.xpath('//*[@class="totalPage"]')  # 获取到章节最大页码
            # pn_high.append(url_f[0].text)
            pnp = url_f[0].text

            for pn in range(1, int(pnp) + 1):

                # 加key不加.webp
                urlz = 'https:{}2F{}%E8%AF%9D{}%2F{}.jpg-zymk.middle'.format(url_tou, ZhangJie + 1 - i, url_key, pn)
                # 加key 加.webp
                urlz1 = 'https:{}2F{}%E8%AF%9D{}%2F{}.jpg-zymk.middle.webp'.format(url_tou, ZhangJie + 1 - i, url_key,pn)
                # 不加key 加.webp
                urlz2 = 'https:{}2F{}%E8%AF%9D{}%2F{}.jpg-zymk.middle.webp'.format(url_tou, ZhangJie + 1 - i, '', pn)
                # 不加key 不加.webp
                url_s = 'https:{}2F{}%E8%AF%9D%2F{}.jpg-zymk.middle'.format(url_tou, ZhangJie, pn)

                # print(ZhangJie + 1 - i, url_key)
                # print(urlz)
                req_baocun = requests.get(urlz)

                if req_baocun.status_code == 200:
                    # print(url)

                    with open('{}\第{}话-{}节.jpg'.format(book_names, ZhangJie + 1 - i, pn), 'wb') as f:
                        f.write(req_baocun.content)
                        # print('第{}话-{}节-保存完成'.format(ZhangJie+1-i, pn))
                        confirmLabel.insert(END, '{}第{}话-第{}节-保存完成'.format(book_names, ZhangJie + 1 - i, pn))
                        confirmLabel.see(END)  # 光标移动到最后显示
                        window.update()  # 刷新文本框显示内容

                elif requests.get(urlz1).status_code == 200:

                    req_baocun = requests.get(urlz1)
                    # print(urlz1)
                    with open('{}\第{}话-{}节.jpg'.format(book_names, ZhangJie + 1 - i, pn), 'wb') as f:
                        f.write(req_baocun.content)
                        # print('第{}话-{}节-保存完成'.format(ZhangJie+1-i, pn))
                        confirmLabel.insert(END, '{}第{}话-第{}节-保存完成'.format(book_names, ZhangJie + 1 - i, pn))
                        confirmLabel.see(END)  # 光标移动到最后显示
                        window.update()  # 刷新文本框显示内容

                elif requests.get(urlz2).status_code == 200:

                    req_baocun = requests.get(urlz2)
                    # print(urlz2)
                    with open('{}\第{}话-{}节.jpg'.format(book_names, ZhangJie + 1 - i, pn), 'wb') as f:
                        f.write(req_baocun.content)
                        # print('第{}话-{}节-保存完成'.format(ZhangJie+1-i, pn))
                        confirmLabel.insert(END, '{}第{}话-第{}节-保存完成'.format(book_names, ZhangJie + 1 - i, pn))
                        confirmLabel.see(END)  # 光标移动到最后显示
                        window.update()  # 刷新文本框显示内容

                elif requests.get(url_s).status_code == 200:

                    # print("失败,重新拼接URL")
                    req_baocun_s = requests.get(url_s).content
                    print(url_s)
                    with open('{}\第{}话-{}节.jpg'.format(book_names, ZhangJie + 1 - i, pn), 'wb') as f:
                        f.write(req_baocun_s)
                    # print('第{}话-{}节-保存完成'.format(ZhangJie+1-i, pn))
                    confirmLabel.insert(END, '{}第{}话-第{}节-保存完成'.format(book_names, ZhangJie + 1 - i, pn))
                    confirmLabel.see(END)  # 光标移动到最后显示
                    window.update()  # 刷新文本框显示内容
    print('已经全部保存完成')
    confirmLabel.insert(END, '已经全部保存完成')
    confirmLabel.see(END)  # 光标移动到最后显示
    window.update()  # 刷新文本框显示内容


# 查找书籍的id函数
def book_name(namee):
    n = namee
    url = 'https://www.zymk.cn/api/getsortlist/?callback=getsortlistCb&key={}&topnum=20&client=pc'.format(n)
    req = requests.get(url)  # .content
    res = re.findall('"comic_id":\d+', str(req.text))
    res = re.findall('\d+', str(res))
    res1 = re.findall('"comic_name":"\s*?\S*?"', str(req.text))
    res1 = re.findall(':"\s*\S*"', str(res1))
    res2 = re.sub(r'[\/\\\:\*\?"\<\>\|\[\]\.]', '', str(res1))
    res2 = re.findall(r"'(.+?)'", res2)

    zidian = dict(zip(res, res2))
    # print('查找到以下内容：', '\n', zidian)
    return zidian


# 选择要现在的书籍函数
def namee():
    confirmLabel.delete(0, END)  # 清空文本框
    namee = namee_Entry.get()
    zidian = book_name(namee)  # 调用搜索书名函数

    if len(zidian) != 0:  # 返回的字典为空，提示重新输入，否则正常执行
        confirmLabel.insert(END, '\t', '                             请双击要下载的漫画:', '\t')

        for i in zidian.items():
            confirmLabel.insert(END, i)
            # confirmLabel.see(END)  # 光标移动到最后显示

    else:
        # print('请输入要下载的漫画：')
        confirmLabel.insert(END, '请输入要下载的漫画：' + '\t')


def xuanze(event):  # 选择要下载的漫画
    zidian = confirmLabel.get(confirmLabel.curselection())
    # print(zidian)
    if type(zidian) == tuple:  # 判断点击是否为搜索的内容，
        confirmLabel.delete(0, END)  # 清空文本框
        confirmLabel.insert(END, '开始下载:', '\n', zidian[1])

        # print('开始下载 {}'.format(zidian[1]))
        isExists = os.path.exists('./{}'.format(zidian[1]))
        if not isExists:
            os.mkdir(zidian[1])

        # print('https://www.zymk.cn/{}/'.format(zidian[0]), zidian[1])
        all_all('https://www.zymk.cn/{}/'.format(zidian[0]), zidian[1])


window = Tk()
window.geometry('600x600+500+200')  # 窗口大小
window.title('漫画下载--本程序将搜索知音漫客网站信息')

taitouLabel = Label(window, text="请输入要下载的漫画:  ", height=4, width=30, font=("Times", 20, "bold"), fg='red')

namee_Entry = Entry(window, width=25, font=("Times", 20, "bold"))

button = Button(window, text="搜索", command=namee, )  # .grid_location(33,44)
GunDongTiao = Scrollbar(window)  # 设置滑动块组件
confirmLabel = Listbox(window, height=15, width=55, font=("Times", 15, "bold"), fg='red', bg='#EEE5DE',
                       yscrollcommand=GunDongTiao.set)  # Listbox组件添加Scrollbar组件的set()方法
# window.iconbitmap('timg.ico')#设置窗口图标
confirmLabel.bind('<Double-Button-1>', xuanze)  # 双击选择文本框的内容

GunDongTiao.config(command=confirmLabel.yview)  # 设置Scrollbar组件的command选项为该组件的yview()方法

taitouLabel.grid(column=1)
namee_Entry.grid(row=1, column=1, sticky=N + S)
button.grid(row=1, column=1, sticky=E)

confirmLabel.grid(row=3, column=1, sticky=E)
GunDongTiao.grid(row=3, column=2, sticky=N + S + W)  # 设置垂直滚动条显示的位置
window.mainloop()




