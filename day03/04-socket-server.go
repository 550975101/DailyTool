package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", address)
	//简写 冒号前面默认是本机 localhost
	//net.Listen("tcp", ":8848")
	if err != nil {
		fmt.Println("net listener err:", err)
		return
	}

	fmt.Println("监听中....")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener accept err:", err)
		return
	}
	fmt.Println("连接建立成功!")
	//创建一个容器,用于接收读取的数据
	//使用make来创建字节切片 byte==> unit8
	buf := make([]byte, 1024)

	//cnt: 真正读取client发来的数据的长度
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	fmt.Println("client ====>  server,长度: ", cnt, " ,数据: ", string(buf))

	//服务器对客户端请求进行响应
	//将数据转成大写 "hello" ==》 HELLO
	//func ToLower(s string) string
	upperData := strings.ToUpper(string(buf))

	//Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	if err != nil {
		return
	}
	fmt.Println("Client <==== server,长度: ", cnt, ",数据: ", upperData)

	//关闭链接
	err = conn.Close()
	if err != nil {
		fmt.Println("关闭失败: ", err)
		return
	}
}
