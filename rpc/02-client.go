package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1、用rpc链接服务器--Dial()
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("rpc.Dial err: ", err)
		return
	}
	defer conn.Close()

	//2、调用远程函数
	//接收返回值  传出参数
	var reply string
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("Call: err", err)
		return
	}
	fmt.Println(reply)
}
