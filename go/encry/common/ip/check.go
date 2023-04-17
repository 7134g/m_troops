package ip

import (
	"encry/common/logs"
	"encry/config"
	"net"
	"strconv"
	"strings"
)

func CheckPortExist(addr string) {
	conn, _ := net.Dial("tcp", addr)
	if conn == nil {
		return
	}
	defer conn.Close()

	// 占用
	logs.Error("The connection port is in conflict, port:", addr)
}

func CheckPort(port string) string {
	PortNum, err := strconv.Atoi(port)
	if err != nil {
		logs.Error("port should be a number")
	}
	if PortNum < 1 || PortNum > 65535 {
		logs.Error("port should be a number and the range is [1,65536)")
	}
	addr := "127.0.0.1:" + port
	CheckPortExist(addr)

	return port
}

func CheckIp(address string) bool {
	_, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		logs.Error("Incorrect incoming address: " + address)
	}

	return true
}

func checkKey(index *int, args []string) {

	if *index+1 < len(args) {
		if exist := strings.Contains(args[*index+1], "-"); !exist {
			*index++
			//CONFUSE = checkKey(args[index])
			lkey := len(args[*index])
			if lkey != 16 {
				var patch []string
				if lkey < 16 {
					for i := 0; i < 16-lkey; i++ {
						patch = append(patch, strconv.Itoa(0))
					}
					config.CONFUSE = args[*index] + strings.Join(patch, "")
				} else {
					config.CONFUSE = args[*index][:16]
				}
			} else {
				config.CONFUSE = args[*index]
			}
		}
		logs.Info("key set successfully")
	}
	return
}
