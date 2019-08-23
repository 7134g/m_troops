import re
import requests
from fake_useragent import UserAgent
import os


PAGE = 999
VALID = 0
INVALID = 0

#设置自动化headers对象
UA = UserAgent()


def test_url(http):
    url = 'https://www.baidu.com/'
    headers = {
                'User_Agent':UA.random,
                'Host':'www.baidu.com',
                'Connection':'keep-alive',
                'Cache-Control':'max-age=0',
                'Accept':'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8'
                }
    proxies = {'http':http,
               'https':http}
##    print(headers)
##    print(proxies)
    res = requests.get(url,headers=headers,proxies=proxies,timeout=5)

    return res

def ip_url(url):
    headers = {'User-Agent':UA.random}
    html = requests.get(url,headers=headers)

    return html

#获取每一页的ip地址，一共 15 * x 个
def get_onepage_ip(PAGE):
    os.chdir(os.getcwd())
    for page in range(1,PAGE):
        url = 'https://www.kuaidaili.com/free/inha/{}/'.format(page)
        html = ip_url(url)
        #ip地址
        p = re.compile('(?:(?:[0,1]?\d?\d|2[0-4]\d|25[0-5])\.){3}(?:[0,1]?\d?\d|2[0-4]\d|25[0-5])')
        alist = p.findall(html.text)
        
        #获取端口
        temp = re.findall('data-title="PORT">\d+',html.text)
        temp = re.findall('\d+',str(temp))
        
        dabao = dict(zip(alist,temp))
        iplist = []
        for i,k in dabao.items():
            iplist.append('http://'+i+':'+k)
        print('开始筛选第{0}批IP,当前一共{1}个IP等待筛选'.format(page,len(iplist)))
        yield iplist

def test_ip_to_save(iplist):

    global VALID,INVALID
    
    #测试爬取的ip是否有效
    file = open('ip.txt','a')
    for http in iplist:
        try:
            res = test_url(http)
            status = res.status_code
            print('网页相应码：'+str(status))
            
            if status == 200:
                VALID += 1
                print('第{0}个有效ip为：{1}'.format(str(VALID),http))
                file.write(http)
                file.write('\n')
            
        except Exception as e:
            INVALID += 1
            print(e)
            print('第{0}个失效的IP，该IP为：{1}'.format(str(INVALID),http))
            
        finally:
            print('---------------------------------------')
            
    print('当前有效的IP现在一共有{}个'.format(str(VALID)))
    file.close()
    print('===============我是下一批ip的分割线==================')
        
def main():    
    # 执行主要程序
    for iplist in get_onepage_ip(PAGE):
        test_ip_to_save(iplist)
        

if __name__ == '__main__':  
    #如果当前目录存在ip.txt文件，则删除
    file = os.getcwd() + '\ip.txt'
    if os.path.exists(file):
        print('检测到存在ip.txt，执行删除')
        os.remove(file)

    main()
    print(input('请按回车键结束'))
