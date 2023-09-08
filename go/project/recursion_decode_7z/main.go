package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

// D:\software\7-Zip\7z.exe x -p 大学.7z
// var num = []string{"1", "2", "3", "4", "5", "6"}
// var num = []string{"0","1","2","3","4","5","6","7","8","9",
//
//		"a","b","c","d","e","f","g","h","i","j","k","l","m","n",
//		"o","p","q","r","s","t","u","v","w","x","y","z",
//		"A","B","C","D","E","F","G","H","I","J","K","L","M","N",
//		"O","P","Q","R","S","T","U","V","W","X","Y","Z",
//	}
var characters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.!@#$%^&*()_+-=[]{}|;\':",<>?/`

// var characters = "0123456789"
var lenPW = 6    // 密码最大长度
var lenMinPW = 0 // 密码起始长度
var fileName = `nothing.zip`
var dirName = "nothing"
var passwordChan = make(chan string, 2)

func main() {
	flag.StringVar(&characters, "k", characters, "用来构造密码的内容，默认值："+characters)
	flag.IntVar(&lenPW, "l", lenPW, "默认密码长度为6")
	flag.IntVar(&lenMinPW, "lm", lenMinPW, "默认最短密码长度为0")
	flag.StringVar(&fileName, "n", fileName, "默认为测试文件 nothing.zip")
	flag.Parse()

	if _, err := os.Stat(fileName); err != nil {
		log.Fatal("找不到", fileName)
	}

	go monitor()
	for _, fn := range copyFile() {
		go run(fn)
	}

	generateAllPossibleStrings(lenPW)
}

func copyFile() []string {
	fns := make([]string, 0)
	fs := make([]io.Writer, 0)
	ext := path.Ext(fileName)
	for i := 0; i < 5; i++ {
		fn := strings.ReplaceAll(fileName, ext, "")
		fnPath := filepath.Join("pack", fmt.Sprintf("%s_%d%s", fn, i, ext))
		f, err := os.Create(fnPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fs = append(fs, f)
		fns = append(fns, fnPath)
	}

	w := io.MultiWriter(fs...)
	r, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w, r); err != nil {
		panic(err)
	}

	return fns
}

func generateAllPossibleStrings(lenPW int) {
	for i := 0; i <= lenPW; i++ {
		generate("", characters, i)
	}

}

func generate(prefix string, characters string, remainingLength int) {
	if remainingLength == 0 {
		//*result = append(*result, prefix)
		passwordChan <- prefix
		return
	}

	for _, char := range characters {
		newPrefix := prefix + string(char)
		generate(newPrefix, characters, remainingLength-1)
	}
}

func run(fileName string) {
	dirName = strings.ReplaceAll(fileName, path.Ext(fileName), "")
	_ = os.MkdirAll(dirName, os.ModeDir)

	ticker := time.NewTicker(time.Second * 30)
	for {
		select {
		case pw := <-passwordChan:
			atomic.AddInt64(&count, 1)
			work(fileName, dirName, pw)
			ticker.Reset(time.Second * 5)
		case <-ticker.C:
			fmt.Println("time over")
			os.Exit(0)
		}
	}
}

func work(fileName, dirName, pw string) {
	//fmt.Println(pw)
	ext := path.Ext(fileName)
	switch ext {
	case ".zip":
		if err := DeCompressZip(fileName, dirName, pw, nil, 0); err == nil {
			fmt.Printf("ok ===> %s  count ====> %d\n", pw, count)
			os.Exit(0)
		}
	case ".rar":
		if err := DeCompressRar(fileName, dirName, pw); err == nil {
			fmt.Printf("ok ===> %s  count ====> %d\n", pw, count)
			os.Exit(0)
		}
	default:
		panic("file type err")

	}

}
