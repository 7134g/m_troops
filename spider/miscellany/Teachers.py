# -*- coding: utf-8 -*-
"""
Created on Wed Mar 21 11:27:08 2018

@author: coyota
"""
try :
    import requests
    from PIL import Image
    import time
    import re
    import openpyxl
except Exception as reason:
    print("存在未成功安装的库，请确认安装后重新运行\n")
    print("keyword:",reason)


def get_vrifycode_session():     #获取验证码以及构建session
    Atimer = int((time.time())*1000)
    sessions = requests.Session()
    VerifyCode_url = 'http://jxfw.gdut.edu.cn/yzm?d='+str(Atimer)
    r = sessions.get(VerifyCode_url)
    r= r.content
    with open("v.jpg","wb") as i:
        i.write(r)
    img = Image.open("v.jpg")
    img.show()
    img.close()
    return sessions


def makezcs(kbxx):   #整理课表中上课周次为指定格式
    for item in kbxx :
        zcs = list(eval(item["zcs"]))
        zcs.sort()
        flag = 1
        during = []
        content = []
        for i in range(1,len(zcs)) :
            if zcs[i] == zcs[i-1] + 1 :
                flag += 1
            else :        
                during.append(flag)
                flag = 1
        during.append(flag)
        flag = 0
        for i in during :
            if i == 1:
                content.append(str(zcs[flag])+"周")
            else :
                content.append(str(zcs[flag])+"-"+str(zcs[flag+i-1])+"周")
            flag += i
        for i in content :
            if i == content[0] :
                str1 = str(i)
            else :
                str1 = str1 + "&" + str(i)
        item["zcs"] = str1
    return kbxx


def makejcdm2(kbxx): #整理课表中上课星期、节次为指定格式
    weekname = ["","周一 ","周二 ","周三 ","周四 ","周五 ","周六 ","周日 "]
    
    for item in kbxx :             
        jcdm2 = re.findall("[1-9][0-2]?",item["jcdm2"])
        for each in range(len(jcdm2)) :
            jcdm2[each] = int(jcdm2[each])
        jcdm2.sort()        
        item["jcdm2"] = weekname[int(item["xq"])] + str(jcdm2[0]) + "-" + str(jcdm2[-1]) + "节"    
    return kbxx


def SaveClassSchedule(kbxx): #保存课表中所需数据为excel文件。
    kb1 = makejcdm2(kbxx)
    kbxx = makezcs(kb1)
    wb = openpyxl.Workbook()
    ws = wb.active
    ws.append(["name","type","time","during","teacher","place"])
    for item in kbxx :
        ws.append([item["kcmc"],"必修",item["jcdm2"],item["zcs"],item["teaxms"],item["jxcdmcs"]])
    wb.save("kbxx.xlsx")



def main():
    data={}
    data['account'] = input("账号:")
    data['pwd'] = input("密码:")
    Session = get_vrifycode_session()
    data['verifycode'] = input("验证码:")
    loginurl = "http://jxfw.gdut.edu.cn/new/login"
    r = Session.post(loginurl,data = data)
    kburl = "http://jxfw.gdut.edu.cn/xsgrkbcx!xsAllKbList.action?xnxqdm=201702"
    
    r = Session.get(kburl)
    strkb = re.findall(r"kbxx.+}\]",r.text)
    kbxx = eval(strkb[0].split("=")[1])
    SaveClassSchedule(kbxx)

if __name__ == "__main__" :
    main()
    input("导出成功，回车键退出")
