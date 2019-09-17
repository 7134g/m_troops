#coding=utf-8

import requests
import xml.dom.minidom
import time

def TimeStampToTimeString(timestamp, all_info):
    if all_info:
        return time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(timestamp))
    else:
        return time.strftime('%Y-%m-%d', time.localtime(timestamp))
    

def ParseTime(data):
    list1 = data.split(',')
    video_time = float(list1[0])
    send_time = int(list1[4])
    
    video_time = '%.2d-%.2d-%.2d' % ((video_time / 60 / 60) % 60, (video_time / 60) % 60, video_time % 60)
    send_time = TimeStampToTimeString(send_time, True)
    return {'video_time': video_time, 'send_time': send_time}
    

def ParseData(url):
    r = requests.get(url)
    root = xml.dom.minidom.parseString(r.text).documentElement
    
    list1 = []
    for i in root.getElementsByTagName('d'):
        if i.firstChild is None:    # 会出现 i.firstChild 为 NoneType 的情况，只需简单忽略就好
            continue
        list1.append(dict(ParseTime(i.getAttribute('p')), **{'content': i.firstChild.data}))
    
    return list1
    

def SaveData(filename, data):
    with open(filename, 'wb') as f:
        f.write(str('video_time' + '\t' + 'send_time' + '\t\t' + 'content' + '\n\n').encode('utf-8'))
        for i in data:
            f.write(str(i['video_time'] + '\t' + i['send_time'] + '\t' + i['content'] + '\n').encode('utf-8'))
        
    print(filename + '  ......   Done')
    


if __name__ == '__main__':
    url = 'https://comment.bilibili.com/rolldate,13023635'
    video_id = url.split(',')[-1]
    
    r = requests.get(url)
    if r.status_code == 200:
        for i in r.json():
            list1 = ParseData('https://comment.bilibili.com/dmroll,' + i['timestamp'] + ',' + video_id)
            SaveData(TimeStampToTimeString(int(i['timestamp']), False) + '.txt', list1)

