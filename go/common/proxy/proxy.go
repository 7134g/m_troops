// Package proxy 提供一个监听端口，将数据转发出去（简单来说就是拦截，再发送）
package proxy

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"
)

func Serve() error {
	// tcp 连接，监听 8080 端口
	address := "127.0.0.1:8080"
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			return err
		}
		go handle(client)
	}

}

func handle(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	// 用来存放客户端数据的缓冲区
	var b [1024]byte
	//从客户端获取数据
	n, err := client.Read(b[:])
	if err != nil {
		return
	}
	var method, URL, address string
	// 从客户端数据读入 method，url
	_, _ = fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &URL)
	hostPortURL, err := url.Parse(URL)
	if err != nil {
		return
	}
	// 如果方法是 CONNECT，则为 https 协议
	if method == "CONNECT" {
		address = hostPortURL.Scheme + ":" + hostPortURL.Opaque
	} else { //否则为 http 协议
		address = hostPortURL.Host
		// 如果 host 不带端口，则默认为 80
		if strings.Index(hostPortURL.Host, ":") == -1 { //host 不带端口， 默认 80
			address = hostPortURL.Host + ":80"
		}
	}
	//获得了请求的 host 和 port，向服务端发起 tcp 连接
	server, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	defer server.Close()
	if method == "CONNECT" {
		_, _ = fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else { //如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端
		_, _ = server.Write(b[:n])
	}
	// 不使用下级代理，转发数据
	go io.Copy(server, client)
	io.Copy(client, server)
	return
}

func proxyHandle(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	server, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		return
	}
	defer server.Close()

	go io.Copy(server, client)
	io.Copy(client, server)
	return
}
