package main

import (
	"net"
)

func ConnRW(conn1, conn2 net.Conn) {
	//var buffer bytes.Buffer
	//var dataLength int
	//var head []byte
	var spliceData []byte
	var contextLen int
	var spliceLen int
	var step int
	//var patchLen int
	//var indexs []int

	buffer := make([]byte, 0)
	chunk := make([]byte, BUFMAX)
	//var spliceData []byte
	for {
		n, err := conn1.Read(chunk)
		//fmt.Println(flag)
		//if err == io.EOF && n == 0 {
		//	log.Println("[←] Disconnect at local:[" + conn1.LocalAddr().String() + "] and remote:[" + conn1.RemoteAddr().String() + "]")
		//
		//	break
		//}
		if err != nil {
			log.Println("[←] Disconnect at local:[" + conn1.LocalAddr().String() + "] and remote:[" + conn1.RemoteAddr().String() + "]")
			break
		} else if n > 0 {
			existHeader := checkHeader(chunk[:HEADERLEN])
			if existHeader == NOHEAD && len(buffer) == 0 {
				//// no head
				//head = CreateHead()
				//// 拼接原始数据包
				//buffer = append(buffer, head...)
				//buffer = append(buffer, chunk[:n]...)
				//// 加密
				//buffer = checkAESandGZIP(buffer, flag)
				//// 修改头部
				//buffer = changeHeadLenth(buffer)
				////fmt.Println(len(buffer))
				//conn2.Write(buffer)
				//buffer = make([]byte, 0)
				//continue
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

						contextLen = getHeadLenth(spliceData[CSTART:CEND])

						if contextLen == n {
							// 读完全部
							buffer = append(buffer, spliceData...)
							buffer = checkAESandGZIP(buffer)
							conn2.Write(buffer)
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
							conn2.Write(buffer)
							buffer = make([]byte, 0)
							spliceData = spliceData[contextLen:]
						}

						// 重置缓冲区, 将剩余数据进行下一个循环
						//resetData(&buffer, &spliceData, contextLen)
						continue

					} else {
						// 缓冲区有数据
						// 补充剩余数据

						// 头部不完整
						step = 0
						if len(buffer) < HEADERLEN+1 {
							for i := 0; i < HEADERLEN+1; i++ {
								if len(buffer) == HEADERLEN {
									break
								}
								buffer = append(buffer, spliceData[i])
								step++
							}
						}
						spliceData = spliceData[step:]
						contextLen = getHeadLenth(buffer[CSTART:CEND])
						spliceLen = contextLen - len(buffer)                   // 仍需要拼接的长度
						if spliceLen > BUFMAX || spliceLen > len(spliceData) { // 第二次读取仍然不足
							buffer = append(buffer, spliceData[:]...)
							break
						}
						remainingData := spliceData[:spliceLen]
						//fmt.Println("late spliceData: ", len(spliceData), " remainingData: ", len(remainingData))
						buffer = append(buffer, remainingData...) // 获取到了完整数据

						buffer = checkAESandGZIP(buffer)
						conn2.Write(buffer)
						//log.Println("chunk: ", len(temp), "buffer: ", len(buffer), "spliceData: ", len(spliceData), " spliceLen: ", spliceLen)

						// 重置缓冲区, 将剩余数据进行下一个循环
						buffer = make([]byte, 0)
						//log.Println("late spliceData: ", len(spliceData), " spliceLen: ", spliceLen, " remainingData: ", len(remainingData))
						spliceData = spliceData[spliceLen:]
						//resetData(&buffer, &spliceData, spliceLen)
						continue

					}

				}

			}
		}
	}

	log.Println("[←] Disconnect other at local:[" + conn2.LocalAddr().String() + "] and remote:[" + conn2.RemoteAddr().String() + "]")
	conn2.Close()
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
	conn2.Write(buffer)
}

//func resetData(buffer, spliceData *[]byte, L int) {
//	*spliceData = (*spliceData)[L:] // 将剩余数据进行下一个循环
//	*buffer = make([]byte, 0)
//}
