package main

// OS指令
var (
	// aes混淆密匙
	CONFUSE = "^($*(897@6>8<1?9"
	// 日志是否开启
	LOGTPYE string
	// aes和gzip是否开启
	SIGN []string
	// listen模式参数
	LISTEN []string
	// tran模式参数
	TRAN []string
	// slave模式参数
	SLAVE []string
)

// 头部构造
var HEADER = "*%&$0000!00*"

const (
	// 头部长度
	HEADERLEN = 12
	// 头部文本数据长度起始位置
	CSTART = 4
	// 头部文本数据长度结束位置
	CEND = 8

	// gzip在头部的位置
	IGIP = 9
	// aes在头部的位置
	IAES = 10

	// 缓冲区大小
	BUFMAX = 1024
	// 无头
	NOHEAD = 0
	// 有头
	HAVEHEAD = 1
)

const (
	// 超时时间
	timeout = 3

	// 头部中代表加密操作的取值范围
	MIN_SCOPE = 48
	MAX_SCOPE = 52

	// 原文
	ORIGIN = 48
	// 加密
	ENCODE = 49
	// 解密
	DECODE = 50

	// 同时加密压缩
	GZIPAES = 51
	// 同时解密解压
	UNGZIPAES = 52
)
