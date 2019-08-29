import socket
from selectors import DefaultSelector, EVENT_READ, EVENT_WRITE # 自动采用select或者epoll
from urllib.parse import urlparse
import time


def timer(func):
    def wrapper(*args,**kwargs):
        t1 = time.time()
        v = func(*args,**kwargs)
        t2 = time.time()
        z_time = t2-t1
        print("本次操作耗费{}秒".format(z_time))
        return v
    return wrapper


class Fecher:
    def __init__(self,headers = None):
        self.host = ''
        self.path = ''
        self.data = b''
        # self.client = ''
        self.headers = headers

    def connected(self, key):
        selector.unregister(key.fd)

        request_contain = ''
        if self.headers:
            for k, v in self.headers.items():
                temp = "{k}:{v}\r\n".format(k=k, v=v)
                request_contain = "".join([request_contain, temp])

        send_data = 'GET {path} HTTP/1.1\r\nHost:{host}\r\n{request_contain}\r\n'.format(path=self.path, host=self.host, request_contain=request_contain).encode('utf8')
        print("发起信息", send_data)

        self.client.send(send_data)

        selector.register(self.client.fileno(), EVENT_READ, self.readable)

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

    def get(self, url):
        self.spider_url = url
        url = urlparse(url)
        self.host = url.netloc
        self.path = url.path
        if not self.path:
            self.path = '/'

        self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.client.setblocking(False)  # 非阻塞

        try:
            self.client.connect((self.host, 80))
        except BlockingIOError:
            pass

        # 注册文件描述符
        selector.register(self.client.fileno(), EVENT_WRITE, self.connected)


def creat_loop():
    while not stop:
        ready = selector.select()
        for key, mask in ready:
            call_back = key.data
            call_back(key)


@timer
def main():
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
        "Connection": "close",
    }

    for index in range(10):
        url = 'http://shop.projectsedu.com/goods/{}/'.format(index)
        urls.append(url)
        feture = Fecher(headers=headers)
        feture.get(url)
    creat_loop()


selector = DefaultSelector()
urls = []
stop = False

if __name__ == '__main__':
    main()
