package system

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"golang.org/x/exp/mmap"
	"log"
	"os"
	"runtime"
)

// CheckPermission 检测是否为管理员权限目录下
func CheckPermission(p string) (bool, error) {
	// 检测是否为管理员权限目录下
	f, err := os.Stat(p)
	if err != nil {
		return false, err
	}

	// 判断是否拥有写权限
	modeFile := string(f.Mode().String()[2]) + string(f.Mode().String()[5]) + string(f.Mode().String()[8])
	if modeFile == "www" {
		return true, nil
	}
	return false, nil
}

// GetHardDiskSpace 获取该路径下磁盘容量
func GetHardDiskSpace(s string) float64 {
	var info *disk.UsageStat
	var err error
	if runtime.GOOS == "windows" {
		info, err = disk.Usage(s[:2]) // window
	} else {
		info, err = disk.Usage("/") // linux
	}
	if err != nil || info == nil {
		log.Println("get disk error, info is nil, ", err)
		// 磁盘空间获取失败直接默认剩余10g空间
		return 10240.0
	}

	free := float64(info.Free) / 1024 / 1024
	// mb
	return free
}

// CPUType 获取cpu类型
func CPUType() int {
	hInfo, _ := host.Info()
	if hInfo.KernelArch == "x86_64" {
		return 64
	} else {
		return 32
	}
}

// FileMapping 文件映射
func FileMapping() {
	at, _ := mmap.Open("./tmp.txt")
	buff := make([]byte, 2)
	_, _ = at.ReadAt(buff, 4)
	_ = at.Close()
	fmt.Println(string(buff))
}
