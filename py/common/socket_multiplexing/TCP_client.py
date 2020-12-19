import socket

client = socket.socket()
client.connect(('127.0.0.1', 8000))

while True:
    deal_data = input('我是客户端，请输入发送的数据：')
    client.send(deal_data.encode('utf8'))
    data = str(client.recv(1024), encoding="utf8")
    print("收到服务端发来的数据：{}".format(data))
    if 'close' in data:
        break
client.close()