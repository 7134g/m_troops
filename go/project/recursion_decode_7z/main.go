package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

// D:\software\7-Zip\7z.exe x -p 大学.7z
//var num = []string{"1", "2", "3", "4", "5", "6"}
//var num = []string{"0","1","2","3","4","5","6","7","8","9",
//	"a","b","c","d","e","f","g","h","i","j","k","l","m","n",
//	"o","p","q","r","s","t","u","v","w","x","y","z",
//	"A","B","C","D","E","F","G","H","I","J","K","L","M","N",
//	"O","P","Q","R","S","T","U","V","W","X","Y","Z",
//}
var num = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var stop = false
var dep = 1     // 用于退出递归
const lenPW = 6 // 密码最大长度

func main() {
	Decode("")
}

func Decode(pw string) {
	if stop {
		return
	}
	if dep > lenPW {
		// 解码失败，密码大于预设长度
		return
	}
	dep++
	for _, s := range num {
		password := pw + s
		if Run(password) {
			return
		} else {
			Decode(password)
		}
	}

}

func Run(pw string) bool {
	if stop {
		return false
	}
	fmt.Println(pw)
	cmd := exec.Command("D:\\software\\7-Zip\\7z.exe", "x", "-otemp", "-p"+pw, `大学.7z`)
	if err := cmd.Run(); err != nil {
		RemoveDir()
		return false
	} else {
		fmt.Println("ok ===> ", pw)
		stop = !stop
		return true
	}
}

func RemoveDir() {
	dirs, err := ioutil.ReadDir("temp")
	if err != nil {
		return
	}
	for _, dir := range dirs {
		p := path.Join([]string{"temp", dir.Name()}...)
		_ = os.RemoveAll(p)
	}
}
