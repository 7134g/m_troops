import requests
import re
import os

def baocun_img(zidian,book_name):
        n = 0
        for i, k in zidian.items():
            path = r"C:\Users\Administrator\Desktop\斗破苍穹\{}".format(book_name[n])
            os.makedirs(path)

            for pn_pn in range(1, int(k)+1):
                url = 'https://mhpic.cnmanhua.com/comic/D%2F%E6%96%97%E7%A0%B4%E8%8B%8D%E7%A9%B9%E6%8B%86%E5%88%86%E7%89%88%2F{}%E8%AF%9D%2F{}.jpg-mht.middle'.format(i,pn_pn)
                re = requests.get(url)
                with open(path+'\{}{}.jpg'.format(book_name[n], pn_pn),'wb') as f:
                    f.write(re.content)
                    f.close()
            print(book_name[n],'已经完成')
            n += 1

def book_url():
    url = 'http://www.manhuatai.com/doupocangqiong/'

    res = requests.get(url)
    res.encoding = 'utf-8'
    book_nam = re.findall('title="第\d+话\s\S+\（\S\）', str(res.text))#得到每一话的名字
    book_ = re.findall('第\d+话\s\S+\（\S\）\(\d+P\)', str(res.text))
    book_p =re.findall('\d+P', str(book_))
    book_pn = re.findall('\d+', str(book_p))

    book_name1 = re.sub(r'title="|\s|\（|\）', '', str(book_nam))
    book_name = re.findall(r"'(.+?)'",book_name1)#将字符串转化为list

    book_i = re.findall(r"\d+话", book_name1)
    book_id =re.findall(r'\d+', str(book_i))

    zidian = dict(zip(book_id,book_pn))

    baocun_img(zidian,book_name)


book_url()
