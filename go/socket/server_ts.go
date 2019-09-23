package main

import (
    "fmt"
    "net"
    "strings"
)

var onlineConn = make(map[string]net.Conn)
var messageQ = make(chan string,1000)
var quitChan = make(chan net.Conn)

func main() {
    listen_sock, err := net.Listen("tcp","127.0.0.1:8080")
    CheckError(err)
    defer listen_sock.Close()

    fmt.Println("开始等待请求")

    go consumerMesa()

    for {

        conn, err := listen_sock.Accept()
        CheckError(err)
        fmt.Println("得到一个tcp连接")
        onlineConn[conn.RemoteAddr().String()] = conn
        for name := range onlineConn{
            fmt.Println(name)
        }
        go ProcessInfo(conn)
    }
}

func CheckError(err error)  {
    if err != nil{
        panic(err)
    }
}

// 开始处理连接
func ProcessInfo(conn net.Conn)  {
    buf := make([]byte, 1024)
    addr := conn.RemoteAddr().String()
    defer connClose(conn, addr)

    for {

        readlength, err := conn.Read(buf)
        if err != nil {
            break
        }

        if readlength != 0 {
            mesg := string(buf[:readlength])
            //fmt.Println(mesg)
            messageQ <- mesg

        }

    }

}

func connClose(conn net.Conn, addr string) {
    conn.Close() //关闭连接
    delete(onlineConn, addr) // 释放空间
    fmt.Println("关闭", addr, "连接")
}


// 处理消息
func consumerMesa() {
    for  {
        select {
        case msg := <- messageQ:
            doMesg(msg)
        case conn := <-quitChan:
            quitConn(conn)
        }
    }
}

// 处理消息
func doMesg(msg string)  {
    contant := strings.Split(msg, "#")
    if len(contant) > 1{
        addr := strings.Trim(contant[0]," ")
        message := contant[1]


        // 转发消息
        if conn , ok := onlineConn[addr];ok{
            _, err := conn.Write([]byte(message))
            CheckError(err)
        }else {
        fmt.Println("没有接收对象")
        }

    }else{
        fmt.Println("无法解析消息")
    }
}

func quitConn(conn net.Conn)  {

    addr := conn.RemoteAddr().String()

    conn.Close() //关闭连接
    delete(onlineConn, addr) // 释放空间
    fmt.Println("关闭", addr, "连接")

}