package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	targetPath string
	splitFlag  bool
	mergeFlag  bool
	mergeName  = "merge_file_success"

	splitFileRoot   = "split_file_merge_"
	oneReadDataLen  = 10240      // 10kb
	maxSplitFileLen = 1000000000 // 1gb

	tempDir = "./temp"
)

func main() {
	flag.StringVar(&targetPath, "t", targetPath, "文件路径")
	flag.BoolVar(&splitFlag, "s", splitFlag, "切割文件")
	flag.BoolVar(&mergeFlag, "m", mergeFlag, "合并文件夹内所有merge_file_success文件")
	flag.IntVar(&maxSplitFileLen, "max", maxSplitFileLen, "单个文件最大值")
	flag.StringVar(&mergeName, "name", mergeName, "合并文件,指定名称")
	flag.Parse()

	switch {
	case targetPath == "":
		log.Fatal("需要指定操作的目录或文件")
	case splitFlag:
		splitFile(targetPath)
	case mergeFlag:
		mergeFile(targetPath, mergeName)
	}

}

func splitFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	var i = 1
	for {
		wPath := fmt.Sprintf(splitFileRoot+"%d", i)
		fmt.Println(filepath.Abs(wPath))
		wf, err := os.Create(filepath.Join(tempDir, wPath))
		if err != nil {
			log.Fatal(err)
		}
		fileLength := 0
		for fileLength <= maxSplitFileLen {
			b := make([]byte, oneReadDataLen) // 1gb
			rn, err := f.Read(b)
			if err == io.EOF {
				write(wf, b, rn)
				_ = wf.Close()
				return
			}
			if err != nil {
				log.Fatal(err)
			}

			write(wf, b, rn)
			_ = wf.Close()
			fileLength += rn
			i++
		}

	}
}

func mergeFile(path, name string) {
	wf, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(wf *os.File) {
		_ = wf.Close()
	}(wf)

	if err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(info.Name(), splitFileRoot) {
			fmt.Println(path)
			rf, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			for {
				b := make([]byte, oneReadDataLen)
				rn, err := rf.Read(b)
				if err == io.EOF {
					write(wf, b, rn)
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				write(wf, b, rn)
			}

		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func write(wf *os.File, data []byte, readLen int) {
	wn, err := wf.Write(data[:readLen])
	if err != nil {
		log.Fatal(err)
	}
	if wn != readLen {
		log.Fatal("读写长度不一致")
	}
}
