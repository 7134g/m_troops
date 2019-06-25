# -*- coding: utf-8 -*-
from queue import Queue
from urllib import parse,request
from math import ceil
from http import cookiejar
import time
import random
import hashlib
import msvcrt
import os
import gzip
import json
import re
import threading

#登录参数之一
appid = "15000101"

num = 1
queue = Queue()
lock = threading.Lock()

#默认协议头
DefaultHeaders = {
            "Accept":"*/*",
            "User-Agent":"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET4.0C; .NET4.0E)",
            "Accept-Language":"zh-cn",
            "Accept-Encoding":"gzip;deflate",
            "Connection":"keep-alive",
            "Referer":"http://qzone.qq.com"
}

#设置自动处理cookies
cj = cookiejar.LWPCookieJar()
cookies = request.HTTPCookieProcessor(cj)
opener  = request.build_opener(cookies)

#密码输入，cmd命令行下运行显示*号
def pwd_input():
    chars = []
    while True:
        try:
            newChar = msvcrt.getch().decode(encoding="utf-8")
        except:
            return input("【温馨提醒：当前未在cmd命令行下运行，密码输入无法隐藏】:\n")
        if newChar in "\r\n":
             break
        elif newChar == "\b":
             if chars:
                 del chars[-1]
                 msvcrt.putch("\b".encode(encoding="utf-8"))
                 msvcrt.putch( " ".encode(encoding="utf-8"))
                 msvcrt.putch("\b".encode(encoding="utf-8"))
        else:
            chars.append(newChar)
            msvcrt.putch("*".encode(encoding="utf-8"))
    return ("".join(chars) )



#QQ空间GTK算法
def GetGtk(skey):
    HashId = 5381
    skey = skey.strip()
    for i in range(0, len(skey)):
        HashId = HashId + HashId * 32 + ord(skey[i])
    gtk = HashId & 2147483647
    return gtk

#取cookies对应值
def GetCookie(name):
    for cookie in cj:
        if cookie.name == name:
            return cookie.value

#GET访问
def Http(url,charset="utf-8",headers=DefaultHeaders):
    rr = request.Request(url=url, headers=headers)
    with opener.open(rr) as fp:
        if fp.info().get("Content-Encoding") == 'gzip':
            f = gzip.decompress(fp.read())
            res = f.decode(charset,'ignore')
        else:
            res = fp.read().decode(charset,'ignore')
    return res

#POST访问
def Post(url,postdata,charset="utf-8",headers=DefaultHeaders):
    if postdata:
        postdata = parse.urlencode(postdata).encode("utf-8")
    rr = request.Request(url=url,headers=headers,data=postdata)
    with opener.open(rr) as fp:
        if fp.info().get("Content-Encoding") == "gzip":
            f = gzip.decompress(fp.read())
            res = f.decode(charset)
        else:
            res = fp.read().decode(charset)
    return res

#信息菜单
def menu():
    print("*********主菜单*********")
    print("【*】显示信息")
    print("【1】下载自己的相册")
    print("【2】下载QQ好友的相册")
    print("【0】退出程序")
    print("***********************")

#替换图片名字中的非法文件名符号
def replace(filename):
    p = re.compile(r'\\|\/|\:|\*|\?|\<|>|\"|\|')
    return p.sub(r' ', filename)

#线程池管理
class ThreadPoolMgr():
    def __init__(self,work_queue,fold,thread_num=5): #thread_num 线程数量
        self.threads=[]
        self.work_queue=work_queue
        self.fold=fold
        self.init_threadpool(thread_num)

    def init_threadpool(self,thread_num): #创建线程
        for i in range(thread_num):
            self.threads.append(Mythread(self.work_queue,self.fold));

    def wait_allcomplete(self): #等待线程结束
        for item in self.threads:
            if item.isAlive():
                item.join()

#多线程下载
class Mythread(threading.Thread):
    def __init__(self,work_queue,fold):
        threading.Thread.__init__(self)
        self.work_queue=work_queue
        self.fold=fold
        if not os.path.isdir(self.fold): #判断文件夹是否存在，不存在则创建文件夹
            os.makedirs(self.fold)
        if os.path.isdir(self.fold): #如果文件夹存在，则启动线程
            global num   #使用全局变量num
            num = 1
            self.start()

    def run(self):
        global lock
        global num
        while not self.work_queue.empty(): #队列非空时，一直循环
            url = self.work_queue.get() #取出一条数据
            try:
                try:
                    r = request.urlopen(url["url"],timeout=60)  #下载图片，超时为60秒
                except:
                    r = request.urlopen(url["url"],timeout=120)  #如果超时，再次下载，超时为120秒                 
                
                if 'Content-Type' in r.info():
                    fileName = os.path.join(self.fold,replace(url["name"]+"."+r.info()['Content-Type'].split('image/')[1])) #根据查看返回的“Content-Type”来判断图片格式，然后生成保存路径
                    if lock.acquire(): #线程同步
                        print("开始下载第"+str(num)+"张照片")
                        if os.path.exists(fileName):
                            #图片名称若存在，则重命名图片名称
                            fileName = os.path.join(self.fold,replace("重命名_图片_"+str(num)+"."+r.info()['Content-Type'].split('image/')[1]))                       
                        num=num+1
                        lock.release()
                    f = open(fileName, 'wb') 
                    f.write(r.read())
                    f.close()
                    
            except:
                print(url["url"]+"：下载超时！")
        

#通过比较，取真实相册地址列表——QQ空间存放相册地址的url有四种格式，一一访问，哪个url内的相册地址最多，返回该url
def checklist(hostuin):
    print("正在获取相册列表……")
    alist = ['alist','xalist','hzalist','gzalist']
    num = []
    for i in range(4):
        res = Http("http://"+alist[i]+".photo.qq.com/fcgi-bin/fcg_list_album_v2?inCharset=gbk&outCharset=gbk&hostUin=%s&notice=0&callbackFun=&format=jsonp&plat=qzone&source=qzone&appid=4&uin=%s&t=%s&g_tk=%s"%(hostuin,uin,random.Random().random(),gtk),"GBK")
        if res.find("没有权限") != -1:
            return "没有权限"
            break
        else:
            j = 0
            for match in re.findall('"pre"(.*?)",',res):
                j+=1
            num.append(j)
    for x in range(4):
        if num[x] == max(num):
            return alist[x]
            break

#获取照片地址列表
def getphotolist(hostuin,alist,plist,idcnum):
    res = Http("http://"+alist+".photo.qq.com/fcgi-bin/fcg_list_album_v2?inCharset=gbk&outCharset=gbk&hostUin=%s&notice=0&callbackFun=&format=jsonp&plat=qzone&source=qzone&appid=4&uin=%s&t=%s&g_tk=%s"%(hostuin,uin,random.Random().random(),gtk),"GBK")
    match = re.search('Callback\((.*?)\);',res,re.S) #加re.S 匹配多行
    if match:
        #创建json
        albumJSON = json.loads(match.group(1))
        album = albumJSON["data"]["album"]
        while True:
            print("########相册列表########")
            print("0. 按'0'返回主菜单")
            for i in range (len(album)):
                if album[i]["priv"] == 5:
                    print("%s.《%s》【加密相册】共:%s张"%(i+1,album[i]["name"],album[i]["total"]))
                else:
                    print("%s.《%s》 共:%s张"%(i+1,album[i]["name"],album[i]["total"]))
            print("#######################")
            num = input("输入相册编号进行下载：").strip()
            if num != "0":
                if int(num) < len(album)+1:
                    num = int(num)-1
                    if album[num]["total"] != 0:
                        allphotolist = []
                        for x in range(ceil(album[num]["total"]/100)):
                            pageStart =  x*100
                            res = Http("http://"+plist+".photo.qq.com/fcgi-bin/cgi_list_photo?g_tk=%s&callback=shine0_Callback&mode=0&idcNum=%s&hostUin=%s&topicId=%s&noTopic=0&uin=%s&pageStart=%s&pageNum=100&singleurl=1&notice=0&appid=4&inCharset=gbk&outCharset=gbk&source=qzone&plat=qzone&outstyle=json&format=jsonp&json_esc=1&callbackFun=shine0&_=%s"%(gtk,idcnum,hostuin,album[num]["id"],uin,pageStart,int(time.time()*1000)),"GBK")
                            #修复BUG，替换图片地址中的psbe？为psb？就能正常显示了
                            p = re.compile(r'psbe\?')
                            res = p.sub(r'psb?', res)
                            if res.find("对不起，回答错误") != -1:
                                match = re.search('Callback\((.*?)\);',res,re.S) #加re.S 匹配多行
                                if match:
                                    if plist == "plist":
                                        p = ""
                                    elif plist == "xaplist":
                                        p = "xa."
                                    elif plist == "hzplist":
                                        p = "hz."
                                    elif plist == "gzplist":
                                        p = "gz."
                                    s = json.loads(match.group(1))
                                    question = s["data"]["question"]
                                    answer = input("访问此相册需要回答下面的问题：\n主人提问："+question+"\n你的回答：")
                                    answer = parse.quote(answer)
                                    answer = hashlib.md5(answer.encode("gb2312")).hexdigest().upper()
                                    res = Http("http://"+p+"photo.qq.com/cgi-bin/common/cgi_view_album_v2?inCharset=gbk&outCharset=gbk&hostUin=%s&notice=0&callbackFun=&format=jsonp&plat=qzone&source=qzone&appid=4&uin=%s&albumId=%s&singleUrl=1&t=%s&verifycode=&question=%s&answer=%s&g_tk=%s"%(hostuin,uin,album[num]["id"],int(time.time()*1000),parse.quote(question),answer,gtk),"GBK")
                                    if res.find("对不起，回答错误") != -1:
                                        print("提示：————————————【对不起，回答错误!】")
                                        break
                                    else:
                                        print("提示：————————————【回答正确，请重新选择该相册进行下载!】")
                                        break

                            else:
                                match = re.search('Callback\((.*?)\);',res,re.S)
                                if match:
                                    res = match.group(1)
                                    s = json.loads(res)
                                    photolist = s["data"]["photoList"]
                                    allphotolist.extend(photolist)

                        if len(allphotolist)>0:
                            fold = os.path.join(hostuin,replace(album[num]["name"]).strip())
                            for j in range(len(allphotolist)):
                                queue.put(allphotolist[j])
                            thread = ThreadPoolMgr(queue,fold)
                            thread.wait_allcomplete()

                    else:
                        print("该相册没有照片，请重新输入!")
                else:
                    print("没有该相册编号，请重新输入!")
            else:
                return 0
                break

#初始化取自己相册
def myphoto():
    alist = checklist(uin)
    if alist == "alist":
        plist = "plist"
        idcnum = 0
    elif alist == "xalist":
        plist = "xaplist"
        idcnum = 1
    elif alist == "hzalist":
        plist = "hzplist"
        idcnum = 2
    elif alist == "gzalist":
        plist = "gzplist"
        idcnum = 3
    getphotolist(uin,alist,plist,idcnum)

#初始化取好友相册
def youphoto():
    hostuin = input("请输入好友QQ号码:\n").strip()
    alist = checklist(hostuin)
    if alist != "没有权限":
        if alist == "alist":
            plist = "plist"
            idcnum = 0
        elif alist == "xalist":
            plist = "xaplist"
            idcnum = 1
        elif alist == "hzalist":
            plist = "hzplist"
            idcnum = 2
        elif alist == "gzalist":
            plist = "gzplist"
            idcnum = 3
        getphotolist(hostuin,alist,plist,idcnum)
    else:
        print("对不起，您没有权限进行此操作!")

#QQ登录
class qqlogin:
    #QQ登录加密算法
    def md5(self,string):
        try:
            string = string.encode("utf-8")
        finally:
            return hashlib.md5(string).hexdigest().upper()

    def hexchar2bin(self,num):
        arry = bytearray()
        for i in range(0, len(num), 2):
            arry.append(int(num[i:i+2],16))
        return arry

    def Getp(self,password,verifycode):
        hashpasswd = self.md5(password)
        I = self.hexchar2bin(hashpasswd)
        H = self.md5(I + bytes(verifycode[2], encoding="ISO-8859-1"))
        G = self.md5(H + verifycode[1].upper())
        return G


    #验证码处理
    def GetVerifyCode(self):
        #判断是否需要验证码
        check = Http("http://check.ptlogin2.qq.com/check?regmaster=&uin=%s&appid=%s&r=%s"%(self.uin,appid,random.Random().random()))
        verify =  eval(check.split("(")[1].split(")")[0])
        verify = list(verify)
        if verify[0] == "1":
            img = "http://captcha.qq.com/getimage?uin=%s&aid=%s&%s"%(self.uin,appid,random.Random().random())
            with open("verify.jpg","wb") as f:
                rr = request.Request(url=img, headers=DefaultHeaders)
                f.write(opener.open(rr).read())
            os.popen("verify.jpg")
            verify[1] = input("需要输入验证码，请输入打开的图片\"verify.jpg\"中的验证码：\n").strip()
        return verify

    #登录
    def Login(self,uid,password,verifycode):
        p = self.Getp(password,verifycode)  #密码加密
        url = "http://ptlogin2.qq.com/login?ptlang=2052&u="+uid+"&p="+p+"&verifycode="+verifycode[1]+"&css=http://imgcache.qq.com/ptcss/b2/qzone/15000101/style.css&mibao_css=m_qzone&aid="+appid+"&u1=http%3A%2F%2Fimgcache.qq.com%2Fqzone%2Fv5%2Floginsucc.html%3Fpara%3Dizone&ptredirect=1&h=1&from_ui=1&dumy=&fp=loginerroralert&action=2-14-13338&g=1&t=1&dummy="
        DefaultHeaders.update({"Referer":url}) #更新Referer
        res = Http(url,"utf-8",DefaultHeaders) #GET登录
        if res.find("登录成功") != -1:    
            tempstr =  eval(res.split("(")[1].split(")")[0]) 
            tempstr = list(tempstr)
            print("\n昵称："+tempstr[5]+"，登录成功！") 
            global checklogin
            checklogin = True
        elif res.find("验证码不正确") != -1:
            print("\n验证码错误，请重新登录")
            res = GetVerifyCode()
            res = Login(uin,password,res)
        elif res.find("帐号或密码不正确，请重新输入") != -1:
            print("\n帐号或密码不正确，请重新输入")
            uin = input("请输入QQ号码:\n").strip()
            print("请输入QQ密码:")
            password = pwd_input().strip()
            res = GetVerifyCode()
            res = Login(uin,password,res)
        return res

    #初始化
    def __init__(self,uin,password):
        self.uin = uin  #账号、密码赋值
        self.password = password
        self.res = self.GetVerifyCode() #获取验证码
        self.Login(self.uin,self.password,self.res) #登录

#程序入口
if __name__ == "__main__":
    print("By：鱼C工作室 — comeheres 博客地址：www.8zata.com")
    uin = input("请输入QQ号码:\n").strip()   #输入QQ账号
    print("请输入QQ密码:")                   #输入密码
    password = pwd_input().strip()
    qqlogin(uin,password)                   #登录到QQ网站



if cj:
    skey = GetCookie("skey")
    if skey:
        gtk = GetGtk(skey)

while checklogin:
    menu()
    n = input("请输入您的选择:")
    if n == "1":
        myphoto()
        continue
    elif n == "2":
        youphoto()
        continue
    elif n == "0":
        select=False
        exit()
    elif n == "*":
        menu()
        continue
    else:
        print("选择项不存在，请重新输入！")
        continue


