package main

import "net/rpc"

// 要求服务端在注册rpc对象是 能让编译器检测出 注册对象是否合法

// 创建接口 在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

// RegisterService 调用该方法时 需要给i传参 参数应该是实现了HelloWorld方法的对象
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}
