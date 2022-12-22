package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	//===> 在使用json编码时 这个编码不参与
	Name string `json:"-"`
	//===> 在使用json编码时 这个字段会编码成subject_name
	Subject string `json:"subject_name"`
	//===> 在使用json编码时 将age转成string类型
	//一定要两个字段:名字,类型 中间不能有空格
	Age int `json:"age,string"`
	//===> 在使用json编码时 如果这个字段是空的 那么忽略掉 不参与编码
	Address string `json:"address,omitempty"`
	//gender 是小写的 小写字母开头的 在json编码时会忽略掉
	gender string
}

func main() {

	t1 := Teacher{
		Name:    "张天师",
		Subject: "风水",
		Age:     10,
		Address: "北京市",
		gender:  "女生",
	}
	jsonStr, err := json.Marshal(&t1)
	if err != nil {
		return
	}
	fmt.Println("jsonStr： ", string(jsonStr))
}
