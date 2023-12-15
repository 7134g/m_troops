package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// D:\software\7-Zip\7z.exe x -p 大学.7z
// var num = []string{"1", "2", "3", "4", "5", "6"}
// var num = []string{"0","1","2","3","4","5","6","7","8","9",
//
//		"a","b","c","d","e","f","g","h","i","j","k","l","m","n",
//		"o","p","q","r","s","t","u","v","w","x","y","z",
//		"A","B","C","D","E","F","G","H","I","J","K","L","M","N",
//		"O","P","Q","R","S","T","U","V","W","X","Y","Z",
//	}
//
// var characters = `abcdefghijklmnopqrstuvwxyz0123456789`
var characters = "1234567"

//var characters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.!@#$%^&*()_+-=[]{}|;\':",<>?/`

var maxLen int      // 密码最大长度
var minLen int      // 密码起始长度
var fileName string // 文件名
var fileType string // 文件类型
var passwordChan = make(chan string, 100)

func main() {
	flag.StringVar(&characters, "k", characters, "用来构造密码的内容，默认值："+characters)
	flag.IntVar(&maxLen, "max", 8, "默认密码长度为6")
	flag.IntVar(&minLen, "min", 0, "默认最短密码长度为0")
	flag.StringVar(&fileName, "f", "nothing.zip", "默认为测试文件 nothing.zip")
	flag.StringVar(&fileType, "t", "zip", "设置解析类型")
	flag.Parse()

	if _, err := os.Stat(fileName); err != nil {
		log.Fatal("找不到", fileName)
	}

	go monitor()                    // 监控次数
	go generateAllPossibleStrings() // 密码生成

	if fileType == "" {
		fileType = filepath.Ext(fileName)
	}

	log.Printf("破解类型：%s, 文件名：%v, 起始长度：%d", fileType, fileName, minLen)
	switch fileType {
	case "zip":
		dz, err := NewDeCompressZip(fileName, "./")
		if err != nil {
			log.Fatal(err)
		}
		dz.SetPasswdTask(passwordChan, 5)
		if err := dz.run(); err != nil {
			log.Fatal(err)
		}
	case "rar":
	case "7z":
	default:
		log.Fatal("unknown file type")
	}

}
