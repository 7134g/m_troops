import socket
import threading
import asyncio
import platform

class Server:
    def __init__(self):
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server.setblocking(False)
        self.server.bind(('0.0.0.0', 8000))
        self.server.listen(1)
    
    
    async def start_new_socket(self, sock, addr):
        try:

            while True:
                data = str(sock.recv(1024), encoding="utf8")
                if 'close' in data:
                    sock.send(b'close')
                    break
                print('收到客户端发来的数据：{}'.format(data))
                deal_data = data.upper()
                sock.send(deal_data.encode('utf8'))
    
            print('正常断开{}'.format(threading.currentThread()))
            sock.close()
        except ConnectionResetError:
            print('断开连接')
            pass
    
    def start_loop(self,loop):
        asyncio.set_event_loop(loop)
        loop.run_forever()

    def server_sock(self):
        print("开始监听")
        if 'Win' in str(platform.architecture()):
            loop = asyncio.ProactorEventLoop()
        else:
            import uvloop
            loop = uvloop.new_event_loop()
        threading.Thread(target=self.start_loop, args=(loop,)).start()
        while True:
            sock, addr = self.server.accept()
            print('开始建立连接')
            asyncio.run_coroutine_threadsafe(self.start_new_socket(sock, addr), loop)



if __name__ == '__main__':
    server = Server()
    server.server_sock()