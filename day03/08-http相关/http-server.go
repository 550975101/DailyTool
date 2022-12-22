package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由 router
	//xxxx/user ===> func1
	//xxxx/name ===> func2
	//xxxx/id ===> func3
	//http://127.0.0.1:8080/user func是回调函数 用于路由的响应 这个回调函数原型是固定的
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//writer : ==> 通过writer将数据返回给客户端
		fmt.Println("用户请求详情:")
		fmt.Println("request: ", request)
		//request : ==> 包含客户端发来的数据
		_, _ = io.WriteString(writer, "这是user请求返回的数据")
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是name请求返回的数据")
	})
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是id请求返回的数据")
	})
	fmt.Println("http server start")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http start failed,err", err)
		return
	}
}
