package setting

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DialCtx(TIMEOUT time.Duration) func(ctx context.Context, network, addr string) (net.Conn, error) {
	d := &net.Dialer{
		Timeout:   TIMEOUT, // 超时时间
		KeepAlive: TIMEOUT, // keepAlive 超时时间
	}
	return d.DialContext
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func GetPath() string {
	PWD, _ := os.Getwd()
	rootPathIndex := strings.Index(PWD, PROJECT_NAME)
	return PWD[:rootPathIndex+len(PROJECT_NAME)]
}

// 获取文件路径
func GetFilePath(filePath string) string {
	dbPath := filepath.Join(PROJECT_ROOT, filePath)
	return dbPath
}

// 获取文件路径
func GetSpiderTaskFilePath(spiderName string) string {
	p := strings.Replace(PROJECT_SPIDER_TASK, "*", spiderName, 1)
	dbPath := filepath.Join(GetFilePath(p))
	return dbPath
}

//创建目录
func MakeDir(path string) error {
	filePaths := strings.SplitN(path, `\`, -1)
	dir := strings.Join(filePaths[:len(filePaths)-1], `\`)
	if !IsExist(dir) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return errors.New(fmt.Sprintf("创建文件夹失败: %s", err))
		}
	}

	// 创建文件
	if !IsExist(path) {
		filePtr, err := os.Create(path)
		if err != nil {
			return errors.New(fmt.Sprintf("Create file failed: %s", err))
		}
		defer filePtr.Close()
	}
	return nil
}

//判断文件或目录是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
