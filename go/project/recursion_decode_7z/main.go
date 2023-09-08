package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync/atomic"
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
var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// var characters = "0123456789"
var lenPW = 6 // 密码最大长度
var fileName = `nothing.zip`
var dirName = "nothing"

func main() {
	flag.StringVar(&characters, "k", characters, "用来构造密码的内容，默认值："+characters)
	flag.IntVar(&lenPW, "l", lenPW, "默认密码长度为6")
	flag.StringVar(&fileName, "n", fileName, "默认为测试文件 nothing.zip")
	flag.Parse()

	if _, err := os.Stat(fileName); err != nil {
		log.Fatal("找不到", fileName)
	}
	dirName = strings.ReplaceAll(fileName, path.Ext(fileName), "")
	_ = os.MkdirAll(dirName, os.ModeDir)

	go monitor()

	generateAllPossibleStrings(lenPW)
}

func generateAllPossibleStrings(lenPW int) {
	for i := 4; i <= lenPW; i++ {
		generate("", characters, i)
	}

}

func generate(prefix string, characters string, remainingLength int) {
	if remainingLength == 0 {
		//*result = append(*result, prefix)
		run(prefix)
		return
	}

	for _, char := range characters {
		newPrefix := prefix + string(char)
		generate(newPrefix, characters, remainingLength-1)
	}
}

func run(pw string) {
	atomic.AddInt64(&count, 1)
	fmt.Println(pw)
	if err := DeCompressZip(fileName, dirName, pw, nil, 0); err == nil {
		fmt.Println("ok ===> ", pw)
		os.Exit(0)
	}
}
