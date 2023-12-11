package main

import (
	"os/exec"
	"strings"
)

func ChangeToDoubleHump(str string) string {
	words := strings.Split(str, "_") // 将字符串按照 "_" 分割成多个单词
	for i, word := range words {
		words[i] = strings.Title(word) // 每个单词的首字母大写
	}

	result := strings.Join(words, "") // 连接所有单词
	return result
}

func FmtFile(wPath string) {
	_ = exec.Command("gofmt", "-l", "-w", wPath).Run()
}
