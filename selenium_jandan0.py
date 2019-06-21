
import re
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import requests
import time
import os
def image_baocun(url):
    bytes = requests.get('http://'+url)#,proxies=proxies
    os.makedirs('煎蛋妹子')
    f = open(r'煎蛋妹子\{}'.format(url[-15:]), 'wb')
  
    f.write(bytes.content)
    f.close()
    print('-----------------{}保存完成---------------'.format(url[-15:]))
    time.sleep(1)
def image_id(webpage):
    zz = '[href=]+"//[^\s]*.jpg"'
    huoqu_2 = re.findall(zz, str(webpage))
    print(huoqu_2)
    huoqu_3 = re.sub(r'href="//|"|\[|\]|"|\'|\'|=|//|,', '', str(huoqu_2))
    for i in huoqu_3.split():
        print(i)
        image_baocun(i)

chrome_options = Options()
chrome_options.add_argument("--headless")
driver = webdriver.Chrome(chrome_options=chrome_options)
for pn in range(1, 37):
    print('==========这是第{}页==========='.format(pn))
    driver.get("http://jandan.net/ooxx/page-{}".format(pn))
    webpage = driver.page_source
    image_id(webpage)
    time.sleep(5)

driver.close()

