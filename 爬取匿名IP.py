import re
import urllib.request
import os
import time

def open_url(url):
    req = urllib.request.Request(url)
    req.add_header('User-Agent','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36')
    page = urllib.request.urlopen(url)
    html = page.read().decode('utf-8')

    return html


def get_ip():
    
    #如果当前目录存在ip.txt文件，则删除
    file = os.getcwd() + '\ip.txt'
    if os.path.exists(file):
        os.remove(file)

    
    for page in range(1,50):
        url = 'https://www.kuaidaili.com/free/inha/{}/'.format(page)
        html = open_url(url)
        #ip地址
        p = re.compile('(?:(?:[0,1]?\d?\d|2[0-4]\d|25[0-5])\.){3}(?:[0,1]?\d?\d|2[0-4]\d|25[0-5])')
        alist = p.findall(html)

        #获取端口
        temp = re.findall('data-title="PORT">\d+',html)
        temp = re.findall('\d+',str(temp))
        dabao = dict(zip(alist,temp))
        iplist = []
        for i,k in dabao.items():
            iplist.append('http://'+i+':'+k)
        os.chdir(os.getcwd())

        with open('ip.txt','a') as flie:
            for each in iplist:
                flie.write(each)
                flie.write('\n')
        time.sleep(1)
        print('已经获取：{}'.format(page*15))
        

if __name__ == '__main__':  
    get_ip()
