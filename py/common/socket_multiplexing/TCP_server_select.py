import socket
import selectors

class Server:
    def __init__(self,host, port):
        self.host, self.port = host, port

    def run(self):
        self.creat_loop()
        self.sock.bind((self.host, self.port))
        self.sock.listen(50)
        self.sock.setblocking(False)
        self.sel.register(self.sock,selectors.EVENT_READ,self.accept)
        while True:
            events = self.sel.select()
            for key,mask in events:
                callback = key.data
                callback(key.fileobj,mask)

    def accept(self,sock,mask):
        conn,addr = sock.accept()
        print ('连接来自于{}'.format(addr))
        conn.setblocking(False)
        self.sel.register(conn,selectors.EVENT_READ,self.read)

    def read(self,conn,mask):
        data = conn.recv(1024)
        if data:
            print ('来自客户端：',data)
            data = data.decode().upper()
            conn.send(data.encode("utf8"))
        else:
            print ('准备关闭连接',conn)
            self.sel.unregister(conn)
            conn.close()

    def creat_loop(self):
        self.sel = selectors.DefaultSelector()
        self.sock = socket.socket()
        self.sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)


if __name__ == '__main__':

    server = Server('127.0.0.1',8000)
    server.run()
