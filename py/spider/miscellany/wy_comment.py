import requests
import json

def get_hot_comments(res):
    comments_json = json.loads(res.text)
    hot_comments = comments_json['hotComments']
    with open('hot_comments.text','w',encoding = 'utf-8')as file:
        for each in hot_comments:
            file.write(each['user']['nickname']+ ': \n\n')
            file.write(each['content'] + '\n')
            file.write('-----------------------\n')


def get_comments(url):
    name_id = url.split('=')[1]
    headers = {
        'User-Agent':'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36',
        'Referer':'https://music.163.com/song?id=27867503'
        }

    params = 'KOaUTtrrbcWK0EN1551CbTZ9F79mopu7UfS+PmibTwMpiJtfE/g4X9k0f/aOoSqHZeQ3dyVdPjRIv3H7+94DXqEPEq/Z2MjDq8LtF1cSDU3okubtczHwBBRcCOP59VonJsUtW1gjXNNruntDCW+vVivhzXzB8Tao7VO//N5glFc/XA78nfn0oFhe4W+FYHK3OLdheJUY9nUU9d1gzL7zx9W9zveWkuceKvMv2dL+bZ4='
    encSecKey = '3fedd22442679350e352bb7ad0a1b507200754b81e6316e19d00f83576d8a2edeb335df7313aaab54070bacb04cf408036df7beb542ed34d9bb9fec339f39f2fcaf07115d121da395a01ac77462fd5c4ade5704622721a85da2b5b02cb56bed2fb86b46a2c87e25f28fd225a59f8557204e1ae392a7acbf382df04789ae26ab3'

    data = {'params':params,
            'encSecKey':encSecKey
        }
    target_url = 'https://music.163.com/weapi/v1/resource/comments/R_SO_4_{}?csrf_token=a3232cca345b026de7c0e806ab93b27f'.format(name_id)
    
    res = requests.post(target_url, headers = headers, data = data)

    return res


def main():
    url = input('请输入链接：')
    res = get_comments(url)
    get_hot_comments(res)

    

if __name__ == '__main__':
    main()
