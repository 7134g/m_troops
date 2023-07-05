package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	_ "embed"
)

const (
	markdown = ".md"
)

var (
	HeadMarkdown   = "\n\n%s %s\n\n"
	formatMarkdown = "- [%s](%s)\n"

	detailsMarkdown    = "<details>\n  <summary>点击展开</summary>\n \n\n"
	detailsEndMarkdown = "\n\n</details>\n"
)

var (
	emptyStatus  = false // 是否还未开始写入文件
	openDetail   = false // 生成折叠版本
	markdownName = "README_gen.md"

	ignoreFileName = ".ignore"
	ignoreFileByte []byte
	ignore         = make([]string, 0)
)

func main() {
	//fmt.Println(filepath.Dir("./"))
	//return

	flag.BoolVar(&openDetail, "z", false, "生成折叠版本")
	flag.StringVar(&markdownName, "n", markdownName, "设置名称")
	flag.Parse()

	loadGitignore()
	f, err := os.Create(markdownName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := map[string]map[string]struct{}{}
	_ = filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, "main.go") {
			fmt.Println()
		}
		if judgeFile(path) {
			//fmt.Println("被忽略的文件", path)
			return nil
		}
		if info.IsDir() {
			m[path] = map[string]struct{}{}
			return nil
		}
		dir, _ := filepath.Split(path)
		if strings.HasSuffix(dir, `\`) {
			dir = dir[:len(dir)-1]
		}
		kv, ok := m[dir]
		if !ok {
			kv = map[string]struct{}{}
		}
		kv[path] = struct{}{}
		m[dir] = kv

		return nil
	})

	sortKey := make([]string, 0)
	for k, _ := range m {
		sortKey = append(sortKey, k)
	}

	sort.Strings(sortKey)

	for _, dir := range sortKey {
		fmt.Println(dir)
		kv := m[dir]
		if len(kv) == 0 {
			continue
		}

		writeHeader(f, dir)
		for fp, _ := range kv {
			_, name := filepath.Split(fp)
			writeContent(f, name, fp)
		}
	}

	return

	//_ = filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
	//	if judgeFile(path) {
	//		fmt.Println("被忽略的文件", path)
	//		return nil
	//	}
	//
	//	// ### header
	//	if info.IsDir() {
	//		writeHeader(f, path)
	//		return nil
	//	}
	//
	//	// - [x](a/a.md) 内容
	//	writeContent(f, info.Name(), path)
	//	return nil
	//})
}

func writeHeader(f *os.File, dir string) {
	length := len(strings.Split(dir, string(filepath.Separator)))
	header := "#"

	if emptyStatus && openDetail {
		writeFile(f, detailsEndMarkdown)
	}

	for i := 0; i < length; i++ {
		header = header + "#"
	}
	title := strings.Split(dir, string(filepath.Separator))
	line := fmt.Sprintf(HeadMarkdown, header, title[len(title)-1])
	writeFile(f, line)
	if openDetail {
		writeFile(f, detailsMarkdown)
	}
}

func writeContent(f *os.File, name, path string) {
	fileName := name[:len(name)-3]
	newPath := filepath.ToSlash(path)

	line := fmt.Sprintf(formatMarkdown, fileName, newPath)
	writeFile(f, line)

}

func writeFile(f *os.File, line string) {
	emptyStatus = true
	_, _ = f.Write([]byte(line))
}

func loadGitignore() {
	ignoreFileByte, _ = os.ReadFile(ignoreFileName)
	fileData := string(ignoreFileByte)
	reg, _ := regexp.Compile(`\s+`)
	lines := reg.Split(fileData, -1)

	for _, line := range lines {
		if line == "" {
			continue
		}

		ignore = append(ignore, filepath.ToSlash(line))
	}
}

func judgeFile(path string) bool {
	path = filepath.ToSlash(path)
	if len(ignoreFileByte) == 0 {
		return false
	}
	for _, v := range ignore {
		if v == path {
			return true
		}

	}

	for _, v := range ignore {
		if !strings.Contains(v, "*") {
			continue
		}

		part := strings.SplitN(v, "/*", 2)
		if strings.HasPrefix(path, part[0]) {
			return true
		}

	}

	for _, v := range ignore {
		if v[0] == '*' {
			ext := filepath.Ext(path)
			if ext == v[1:] {
				return true
			}
		}
	}

	return false
}
