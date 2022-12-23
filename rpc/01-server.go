package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// World 定义类对象
type World struct {
}

// 绑定类方法
func (w *World) HelloWorld(name string, resp *string) error {
	*resp = name + "你好"
	return nil
}
func main() {
	//1、注册RPC服务，绑定对象方法
	//参数1: 服务名 字符串类型
	//参数2: 对应的rpc对象,该对象绑定方法满足如下条件
	//      ①、方法必须是导出的--包外可见 首字母大写
	//      ②、方法必须有两个参数 都是导出类型 内建类型
	//      ③、方法的第二个参数必须是"指针" （传出参数）
	//      ④、方法只有一个error 接口类型的 返回值
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册rpc服务失败!")
		return
	}
	//2、设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听...")
	//3、建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept() err: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功...")
	//4、绑定服务
	rpc.ServeConn(conn)
}
