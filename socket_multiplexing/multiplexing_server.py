import socket
from selectors import DefaultSelector, EVENT_READ, EVENT_WRITE
from urllib.parse import urlparse

from _decorator import timer


class Feture:
    def __init__(self, headers = None):
        self.spider_url = ''
        self.host = ''
        self.path = ''
        self.data = b''
        self.headers = headers
        # self.client = ''

    def readable(self, key):
        temp = self.client.recv(1024)
        if temp:
            self.data += temp
        else:
            selector.unregister(key.fd)
            data = self.data.decode('utf8')
            html_data = data.split('\r\n\r\n')
            print(html_data)
            self.client.close()
            urls.remove(self.spider_url)
            if not urls:
                global stop
                stop = True

    def connector(self, key):
        selector.unregister(key.fd)

        request_contain = ''
        for k, v in self.headers.items():
            temp = "{k}:{v}\r\n".format(k=k, v=v)
            request_contain = "".join([request_contain, temp])
        send_data = 'GET {path} HTTP/1.1\r\nHost:{host}\r\n{request_contain}\r\n'.format(
                    path=self.path, host=self.host, request_contain=request_contain).encode('utf8')
        print(send_data)
        self.client.send(send_data)

        selector.register(self.client.fileno(), EVENT_READ, self.readable)


    def get(self, url):
        self.spider_url = url
        url = urlparse(url)
        self.host = url.netloc
        self.path = url.path
        if not self.path:
            path = '/'

        self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            self.client.connect((self.host, 80))
        except BlockingIOError:
            pass
        # print('连接建立')

        selector.register(self.client.fileno(), EVENT_WRITE, self.connector)


def loop():
    while not stop:
        try:
            ready = selector.select()
            for key, mask in ready:
                call_back = key.data
                call_back(key)
        except OSError:
            break

@timer
def main():
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
        "Connection": "close",
    }
    # session = Request(headers=headers)
    # res = session.get("https://www.baidu.com/")

    for index in range(10):
        url = "http://shop.projectsedu.com/goods/{}/".format(index)
        urls.append(url)
        feture = Feture(headers=headers)
        feture.get(url)
    loop()


selector = DefaultSelector()
stop = False
urls = ["https://www.baidu.com/"]

if __name__ == '__main__':
    main()
