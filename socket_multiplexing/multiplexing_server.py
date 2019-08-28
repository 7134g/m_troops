import socket
from selectors import DefaultSelector
from urllib.parse import urlparse

selector = DefaultSelector()





class Request:
    def __init__(self, headers: dict):
        self.headers = headers

    def get(self, url):
        url = urlparse(url)
        host = url.netloc
        path = url.path
        if not path:
            path = '/'

        client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client.connect((host, 80))
        print('开始发送')

        request_contain = ''
        for k,v in self.headers.items():
            temp = "{k}:{v}\r\n".format(k=k, v=v)
            request_contain = "".join([request_contain, temp])
            # headers = 'User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36\r\n'
            # connection = 'Connection:close\r\n'
        send_data = 'GET {path} HTTP/1.1\r\nHost:{host}\r\n{request_contain}\r\n'.format(
            path=path, host=host, request_contain=request_contain).encode('utf8')
        client.send(send_data)

        data = b''
        while True:
            temp = client.recv(1024)
            if temp:
                data += temp
            else:
                break

        data = data.decode('utf8')
        html_data = data.split('\r\n\r\n')
        print(html_data)
        client.close()




if __name__ == '__main__':
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
        "Connection": "close",
    }
    session = Request(headers=headers)
    res = session.get("https://www.baidu.com/")
