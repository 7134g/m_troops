import socket
import threading

class Server:
    def __init__(self):
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server.bind(('0.0.0.0', 8000))
        self.server.listen(1)
    
    
    def start_new_socket(self, sock, addr):
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
    
    
    def server_sock(self):
        print("开始监听")
        while True:
            sock, addr = self.server.accept()
            print('建立连接')
            client_thead = threading.Thread(target=self.start_new_socket,args=(sock,addr))
            client_thead.start()


if __name__ == '__main__':
    server = Server()
    server.server_sock()