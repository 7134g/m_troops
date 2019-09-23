package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    conn, err := net.Dial("tcp","127.0.0.1:8080")
    CheckError(err)
    defer conn.Close()

    go writer_data(conn)

    buf := make([]byte, 1024)
    for  {
       length, err := conn.Read(buf)
       if err != nil{
           break
       }
       msg := buf[:length]
       fmt.Println("server message : "+string(msg))
    }


}

func CheckError(err error)  {
    if err != nil{
        panic(err)
    }
}



func writer_data(conn net.Conn) {
    fmt.Println("请输入数据：")
    for {
        reader := bufio.NewReader(os.Stdin)
        data, _, _ := reader.ReadLine()
        input := string(data)

        if strings.ToUpper(input) == "EXIT"{
            conn.Close()
            break
        }

        _, err := conn.Write([]byte(input))
        if err != nil{
            conn.Close()
            fmt.Println("程序意外退出"+err.Error())
            break
        }

    }
    defer fmt.Printf("客户端结束")
}