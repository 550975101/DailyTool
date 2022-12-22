package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	fmt.Println("client与server连接建立成功!")
	sendData := []byte("hello world")

	cnt, err := conn.Write(sendData)
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
	fmt.Println("Client ===> Server cnt: ", cnt, " ,data： ", string(sendData))

	//接收服务器返回的数据
	//创建buf 用于接收服务器返回的数据
	buf := make([]byte, 1024)
	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn Read err: ", err)
		return
	}
	fmt.Println("Client <===Server ,cnt: ", cnt, ",data: ", string(buf[0:cnt]))

	err = conn.Close()
	if err != nil {
		fmt.Println("关闭失败！,err: ", err)
		return
	}
}
