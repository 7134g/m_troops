import requests
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
import re
import lxml


path = 'pic5/'


def get_one(url):
    browser = webdriver.Firefox()
    wait = WebDriverWait(browser, 10)
    print('正在爬取....')
    try:
        browser.get(url)
        html = browser.page_source
        if html:
            browser.close()
            return html
    except EOFError:
        return None

def pares_one(html):
    soup = BeautifulSoup(html,'lxml')
    imgs  = soup.select('img')
    list = []
    count = 0
    for img in imgs:
        img_url = re.findall('src="(.*?)"',str(img))
        if not img_url[0][-3] == 'gif':
            if not img_url[0][-3] == 'png':
                if img_url[0][:4] == 'http':
                    print('正在下载:%s第%s张' % (img_url[0], count))
                    count += 1
                    list.append(img_url[0])
    return list
def main():
    source_url = 'http://jandan.net/ooxx/page-'
    num = 1
    for i in range(42,48):
        url = source_url +str(i)+'#comments'
        html = get_one(url)
        list = pares_one(html)
        for each in list:
            response = requests.get(each)
            f = open(path + str(num) + '.jpg', 'wb')
            f.write(response.content)
            f.close()
            num += 1

if __name__ == '__main__':
    main()


