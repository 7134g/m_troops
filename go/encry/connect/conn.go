package connect

import (
	"encry/common/logs"
	"encry/config"
	"log"
	"net"
	"sync"
)

func forward(conn1 net.Conn, conn2 net.Conn) {
	logs.Info("[+] start transmit. [%s] <-> [%s] and [%s] <-> [%s] \n", conn1.LocalAddr().String(), conn1.RemoteAddr().String(), conn2.LocalAddr().String(), conn2.RemoteAddr().String())
	var wg sync.WaitGroup
	// wait tow goroutines
	wg.Add(2)
	go connCopy(conn1, conn2, &wg)
	go connCopy(conn2, conn1, &wg)
	//blocking when the wg is locked
	wg.Wait()
}

func connCopy(conn1 net.Conn, conn2 net.Conn, wg *sync.WaitGroup) {
	defer conn1.Close()

	ConnRW(conn1, conn2)
	logs.Info("[←]", "close the connect at local:["+conn1.LocalAddr().String()+"] and remote:["+conn1.RemoteAddr().String()+"]")
	wg.Done()
}

func ConnRW(conn1, conn2 net.Conn) {

	var spliceData []byte
	var contextLen int
	var spliceLen int
	var step int

	buffer := make([]byte, 0)
	chunk := make([]byte, config.BUFMAX)
	//var spliceData []byte
	for {
		n, err := conn1.Read(chunk)

		if err != nil {
			logs.Info("[←] Disconnect at local:[" + conn1.LocalAddr().String() + "] and remote:[" + conn1.RemoteAddr().String() + "]")
			break
		} else if n > 0 {
			existHeader := checkHeader(chunk[:config.HEADERLEN])
			if existHeader == config.NOHEAD && len(buffer) == 0 {
				// no head
				// 拼接原始数据包

				DealNoHead(chunk, n, conn2)
				buffer = make([]byte, 0)
				continue
			} else {
				// 有head
				spliceData = chunk[:n]

				for {
					// 循环处理读出来的数据
					if len(spliceData) == 0 {
						// 无数据处理了
						break
					}

					n = len(spliceData)
					if len(buffer) == 0 {
						// 读取最后数据时，头部不完整
						if len(spliceData) < 13 {
							buffer = append(buffer, spliceData...)
							break
						}

						contextLen = getHeadLenth(spliceData[config.CSTART:config.CEND])

						if contextLen == n {
							// 读完全部
							buffer = append(buffer, spliceData...)
							buffer = checkAESandGZIP(buffer)
							_, _ = conn2.Write(buffer)
							buffer = make([]byte, 0)
							spliceData = spliceData[contextLen:]
						} else if contextLen > n {
							// 数据长度不完整，需要再次read
							buffer = append(buffer, spliceData...)
							break
						} else {
							// 数据长度过长，多个头
							// 修改剩余长度
							buffer = append(buffer, spliceData[:contextLen]...)
							buffer = checkAESandGZIP(buffer)
							_, _ = conn2.Write(buffer)
							buffer = make([]byte, 0)
							spliceData = spliceData[contextLen:]
						}

						// 重置缓冲区, 将剩余数据进行下一个循环
						continue

					} else {
						// 缓冲区有数据
						// 补充剩余数据

						// 头部不完整
						step = 0
						if len(buffer) < 13 {
							for i := 0; i < 13; i++ {
								if len(buffer) == 12 {
									break
								}
								buffer = append(buffer, spliceData[i])
								step++
							}
						}
						spliceData = spliceData[step:]
						contextLen = getHeadLenth(buffer[config.CSTART:config.CEND])
						spliceLen = contextLen - len(buffer) // 仍需要拼接的长度
						if spliceLen > config.BUFMAX {       // 第二次读取仍然不足
							buffer = append(buffer, spliceData[:]...)
							break
						}
						remainingData := spliceData[:spliceLen]
						buffer = append(buffer, remainingData...) // 获取到了完整数据

						buffer = checkAESandGZIP(buffer)
						_, _ = conn2.Write(buffer)

						// 重置缓冲区, 将剩余数据进行下一个循环
						buffer = make([]byte, 0)
						spliceData = spliceData[spliceLen:]
						continue

					}

				}

			}
		}
	}

	log.Println("[←] Disconnect other at local:[" + conn2.LocalAddr().String() + "] and remote:[" + conn2.RemoteAddr().String() + "]")
	_ = conn2.Close()
}

func DealNoHead(chunk []byte, n int, conn2 net.Conn) {
	var buffer []byte
	// no head
	head := CreateHead()
	// 拼接原始数据包
	buffer = append(buffer, head...)
	buffer = append(buffer, chunk[:n]...)
	// 加密
	buffer = checkAESandGZIP(buffer)
	// 修改头部
	buffer = changeHeadLenth(buffer)
	_, _ = conn2.Write(buffer)
}
