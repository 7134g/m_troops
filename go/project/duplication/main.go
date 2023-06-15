package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed README.md
var describe string

var (
	paramsSrc           string
	paramsDst           string
	paramsRemove        bool
	paramsRemoveKeyword string
	paramsDescribe      bool
)

var (
	srcDict       = map[string]string{}
	duplicateDict = map[string][]string{}
)

func main() {
	//flag.StringVar(&paramsSrc, "s", "E:\\documents\\Go\\_programe\\duplication\\a", "源目录")
	//flag.StringVar(&paramsDst, "d", "E:\\documents\\Go\\_programe\\duplication\\b", "对比目录")
	flag.StringVar(&paramsSrc, "s", "", "源目录")
	flag.StringVar(&paramsDst, "d", "", "对比目录")
	flag.BoolVar(&paramsRemove, "r", false, "是否删除对比目录中相同文件")
	flag.StringVar(&paramsRemoveKeyword, "k", "", "删除时匹配路径中的关键字，匹配成功则删除")
	flag.BoolVar(&paramsDescribe, "describe", false, "描述使用方式")
	flag.Parse()

	if paramsDescribe {
		fmt.Println(describe)
		os.Exit(0)
	}

	if paramsSrc == "" || paramsDst == "" {
		log.Fatalln("缺少参数")
	}

	if _, err := os.Stat(paramsSrc); err != nil {
		log.Fatalln("-s 路径错误，尝试使用双引号包裹路径")
	}

	if _, err := os.Stat(paramsDst); err != nil {
		log.Fatalln("-d 路径错误，尝试使用双引号包裹路径")
	}

	_ = filepath.Walk(paramsSrc, conflict("src"))
	_ = filepath.Walk(paramsDst, conflict("dst"))

	for k, list := range duplicateDict {
		fmt.Println(fmt.Sprintf("[src]        %s", k))
		for _, iv := range list {
			msg := "dup"
			if paramsRemove && strings.Contains(iv, paramsRemoveKeyword) {
				_ = os.Remove(iv)
				msg = "del"
			}
			fmt.Println(fmt.Sprintf("    [%s]    %s", msg, iv))
		}
	}
}

func conflict(part string) func(path string, info fs.FileInfo, err error) error {
	fmt.Println(fmt.Sprintf("============== 开始处理 %s ===========", part))
	return func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		key, err := FileMD5(path)
		if err != nil {
			return err
		}

		if v, ok := srcDict[key]; ok {
			if _, _ok := duplicateDict[v]; _ok {
				duplicateDict[v] = append(duplicateDict[v], path)
			} else {
				duplicateDict[v] = []string{
					path,
				}
			}
			return nil
		}
		srcDict[key] = path
		return nil
	}
}

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
