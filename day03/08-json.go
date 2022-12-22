package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Age  int
	//小写字母开头的 在序列化中会忽略掉
	gender string
}

func main() {
	//在网络中传输的时候 把student结构体 编码成json字符串 传输===> 结构体==> 字符串 ==>编码
	//在接收字符串 需要将字符串转换成结构体 然后操作==> 字符串==》结构体==》解码
	lily := Student{
		Id:   1,
		Name: "Lily",
		Age:  20,
		//小写字母开头的 在序列化中会忽略掉
		gender: "女士",
	}
	//编码（序列化），结构==>字符串
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		return
	}
	fmt.Println("encodeInfo: ", string(encodeInfo))

	//对端接收数据
	//反序列化(解码): 字符串==》结构体
	lily2 := Student{}

	err = json.Unmarshal([]byte(encodeInfo), &lily2)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return
	}
	fmt.Println("lily2: ", lily2)
}
