package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}

	response, err := client.Get("http://www.baidu.com")
	if err != nil {
		return
	}

	//beego  gin==> web框架
	ct := response.Header.Get("Content-Type")
	date := response.Header.Get("Date")
	//BWS是百度  Web Server是百度开发的一个web服务器 大部分百度的web应用程序使用的是BWS
	server := response.Header.Get("Server")

	fmt.Println("Content-Type: ", ct)
	fmt.Println("date: ", date)
	fmt.Println("server: ", server)

	url := response.Request.URL
	code := response.StatusCode
	status := response.Status

	fmt.Println("url: ", url)
	fmt.Println("code: ", code)
	fmt.Println("status: ", status)

	body := response.Body
	fmt.Println("body 111", body)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
	}
	fmt.Println("body string : ", string(readBodyStr))
}
