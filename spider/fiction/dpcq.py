import requests
import time
import re

headers={'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.7 Safari/537.36'}
num = 1

def get_info(url):
    res=requests.get(url,headers=headers)
    if res.status_code==200:
        contents = re.findall('<p>(.*?)</p>',res.content.decode('utf-8'),re.S)
        contents = str(contents).replace('&rdquo','')
        contents = str(contents).replace('&hellip','')
        contents = str(contents).replace('&ldquo','')
        contents = str(contents).replace(''''天才一秒记住本站网站 www.doupoxs.com 中间是<span style="color:blue">斗破 拼音+小说 首字母</span> 连起来就是斗破小说，喜欢我就记住我吧！''','')
        
        with open('F:/doupo.txt','a',encoding='utf-8') as f:
            f.write(contents+'\n')
    else:
        pass

if __name__=='__main__':
    urls=['http://www.doupoxs.com/doupocangqiong/{}.html'.format(str(i)) for i in range(2,1665)]
    for url in urls:
        get_info(url)
        print(num)
        num += 1
        time.sleep(1)


