import requests
import re
import os
import time

def baocun_img(zidian,book_name):
        n = 0
        for i, k in zidian.items():
            path = r"C:\Users\Administrator\Desktop\星海镖师\{}".format(book_name[n])
            os.makedirs(path)

            for pn_pn in range(1, int(k)+1):
                url = 'https://mhpic.jumanhua.com/comic/X%2F%E6%98%9F%E6%B5%B7%E9%95%96%E5%B8%88%E6%9D%A1%E6%BC%AB%E7%89%88%2F{}%E8%AF%9DGQ%2F{}.jpg-mht.middle'.format(i,pn_pn)
                re = requests.get(url)
                with open(path+'\{}{}.jpg'.format(book_name[n], pn_pn),'wb') as f:
                    f.write(re.content)
                    f.close()
            print(book_name[n],'已经完成')
            n += 1
            time.sleep(1)

def book_url():
    url = 'https://www.manhuatai.com/xinghaibiaoshi/'

    res = requests.get(url)
    res.encoding = 'utf-8'
    book_nam = re.findall('title="第\d+话\s\S+"', str(res.text))#得到每一话的名字
    book_ = re.findall('第\d+话\s\S+\(\d+P\)', str(res.text))
    book_p =re.findall('\d+P', str(book_))
    book_pn = re.findall('\d+', str(book_p))

    book_name1 = re.sub(r'title="|\s|\（|\）|"', '', str(book_nam))
    book_name = re.findall(r"'(.+?)'",book_name1)#将字符串转化为list

    book_i = re.findall(r"\d+话", book_name1)
    book_id =re.findall(r'\d+', str(book_i))

    zidian = dict(zip(book_id,book_pn))

    baocun_img(zidian,book_name)


if __name__ == '__main__':     
        book_url()

