#################  静态cookie ######################

# class transCookie:
#     def __init__(self, cookie):
#         self.cookie = cookie
#
#     def stringToDict(self):
#         '''
#         将从浏览器上Copy来的cookie字符串转化为Scrapy能使用的Dict
#         :return:
#         '''
#         itemDict = {}
#         items = self.cookie.split(';')
#         for item in items:
#             key = item.split('=')[0].replace(' ', '')
#             value = item.split('=')[1]
#             itemDict[key] = value
#         return itemDict
#
# COOKIE = transCookie('ipLoc-djd=1-72-2799-0; __jdu=1771569446; shshshfpa=fa7f6f93-4129-1880-c2eb-ccc923f1949c-1547380366; shshshfpb=kJLI6C8Rw1incF35IhUjknw%3D%3D; ipLocation=%u5317%u4EAC; areaId=1; mt_xid=V2_52007VwMUVF1cVVIaTB5sAmcAFgFUD1dGGh0bXBliCkBaQQhUX0pVGgxRNQsSWgpYBVsZeRpdBW8fE1JBWFtLH0wSXgFsBxRiX2hSahxOG1oMZQcVUW1YV1wY; __jda=122270672.1771569446.1545053256.1548602424.1549023793.10; __jdv=122270672|direct|-|none|-|1549023793369; o2Control=webp|lastvisit=1; PCSYCityID=1930; shshshfp=bf6e12f8add21e1f7e6989ea3f027d1f; __jdb=122270672.2.1771569446|10.1549023793; __jdc=122270672; shshshsID=384bc0fa98fd39a87449461eea42c509_2_1549024847713')
# COOKIE_CLECK = COOKIE.stringToDict()


########### 动态cookie ##############

import requests
import http.cookiejar as cookielib

import re
class Cookies_cache():

    def __init__(self,accout,password):
        super(Cookies_cache, self).__init__()
        # 初始化请求参数
        self._xsrf = ''
        self.accout = accout
        self.password = password

        self.session = requests.session()


    def get_xsrf(self):
        # 获取_xsrf值
        return self._xsrf


    def login(self):

        # 使用手机登录
        if re.match(r'1\d{10}',self.accout):
            login_url = '/phone'
            print('此时为手机登录')
            data = {
                '_xsrt' : self.get_xsrf(),
                'accout' : self.accout,
                'password':self.password
            }
        elif '@' in self.accout:
            login_url = '/email'
            print('邮箱登录')
            data = {
                '_xsrf':self.get_xsrf(),
                'email':self.accout,
                'password':self.password
            }

        response = self.session.post(login_url,data=data)

        # 将cookies保存在本地
        self.session.cookies = cookielib.LWPCookieJar(filename='cookies.txt')
        self.session.cookies.save()

    # 判断用户是否已经登录
    def is_login(self):
        loading_url = ''
        response = self.session.get(loading_url,allow_redirects=False)
        if response.status_code != 200:
            return False
        else:
            return True