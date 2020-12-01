package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HandError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func download(ch chan int) {
	resp, err := http.Get("http://placekitten.com/g/200/200")
	defer resp.Body.Close()

	HandError(err, "http.Get('http://placekitten.com/g/200/200')")
	imageByte, _ := ioutil.ReadAll(resp.Body)

	filname := "D:\\_document\\Go\\_programe\\test1\\main\\download\\" + strconv.Itoa(int(time.Now().UnixNano())) + ".jpg"
	err = ioutil.WriteFile(filname, imageByte, 0644)
	ch <- 1
	if err == nil {
		fmt.Println(filname, "下载成功！")
	} else {
		fmt.Println(filname, "下载失败！")
	}
}

func main() {
	chs := make([]chan int, 5)
	for n := 0; n < 5; n++ {
		chs[n] = make(chan int)
		go download(chs[n])
	}

	for _, ch := range chs {
		<-ch
	}
}
