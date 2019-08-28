import requests
from urllib.parse import urlencode
from pyquery import PyQuery as pq


base_url = 'https://m.weibo.cn/api/container/getIndex?containerid=2304131677991972&page_type=03&page=2'
fans_url = 'https://m.weibo.cn/api/container/getIndex?containerid=231051_-_fans_-_1677991972&since_id=2'

headers = {
'MWeibo-Pwa': '1',
'Referer':'https://m.weibo.cn/p/2304131663072851_-_WEIBO_SECOND_PROFILE_WEIBO',
'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36',
'X-Requested-With': 'XMLHttpRequest'
}

def get_one_page(page):
    parm1 = {
        'containerid':'2304132360812967',
        'page_type':'03',
        'page':page
    }

    url = base_url + urlencode(parm1)
    try:
        print("start")
        response = requests.get(url,headers=headers)
        print("end")
        if response.status_code == 200:
            return response.json()
        return None
    except requests.ConnectionError as e:
        print('Erro',e.args)
def get_fans_page(sinceid):
    parm2 = {
        'containerid': '2304132360812967',
        'since_id': sinceid
    }
    url = fans_url + urlencode(parm2)
    try:
        response = requests.get(url, headers=headers)
        if response.status_code == 200:
            return response.json()
        return None
    except requests.ConnectionError as e:
        print('Erro', e.args)

def parse_page(json):
    if json:
        items = json.get('data').get('cards')
        for item in items :
            item = item.get('mblog')
            weibo = {}
            weibo['ID'] = item.get('id')
            weibo['内容']= pq(item.get('text')).text()
            weibo['点赞'] = item.get('attitudes_count')
            weibo['评论'] = item.get('comments_count')
            weibo['转发'] = item.get('reposts_count')
            yield weibo
def parse_fans(json):
    if json:
        items = json.get("data").get("cards")
        for item in items:
            card_group = item.get("card_group")
            for card in card_group:
                fans = {}
                fans['ID'] = card.get("itemid")
                fans['scheme'] = card.get("scheme")
                print(fans)



    # if json:
    #     items = json.get('data').get('cards')
    #     for a in items:
    #         fans = {}
    #         fans['ID'] = a.get('buttons')
    #         fans['昵称'] = a.get('screen_name')
    #         fans['他的粉丝数'] = a.get('desc2')
    #         yield fans



def write_to_file(content):
    with open('results1.txt','a',encoding='utf-8') as f:
        f.write(str(content) + '\n')
        f.close()
if __name__=='__main__':
    for sinceid in range(1,100):
        json = get_one_page(sinceid)
        results = parse_fans(json)
        for result in results:
            print(result)
            write_to_file(result)
    for page in range(1,100):
        json = get_one_page(page)
        results = parse_page(json)
        for result in results:
            print(result)
            write_to_file(result)

