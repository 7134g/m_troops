package main

import (
	"encry/common/ip"
	"encry/common/logs"
	"encry/config"
	"os"
)

func main() {
	// 欢迎语
	welcome()

	// 加载配置
	args := os.Args
	argLen := len(os.Args)
	if argLen == 0 {
		loadYaml() // 通过yaml配置启动
	} else {
		loadOS(args, argLen) // 通过命令参数启动
	}

	// 检查
	check()

	// 初始化日志
	logs.Load()

	// 启动
	Run()
}

func check() {
	// 验证log
	if config.LOCALPORT == "" || config.REMOTEADDRESS == "" {
		logs.Exit("args is error")
		return
	}

	// s和c
	if config.StartServer == 1 && config.StartClient == 1 {
		logs.Exit("There are both S and C")
		return
	}

	// 检查是否赋值成功
	if config.REMOTEADDRESS == "" || config.LOCALPORT == "" {
		logs.Exit("remote or local is nil")
		return
	}

	ip.CheckIp(config.REMOTEADDRESS)
	ip.CheckPort(config.LOCALPORT)

}
