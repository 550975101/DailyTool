package main

import (
	"fmt"
	"os"
)

func main() {
	//C: argc, **argv
	//GO: os.Args ==> 直接可以获取命令行输入,是一个字符串切片
	//go build 01-switch.go生成exe
	//执行  ./01-switch.exe hello world test
	//输出 如下
	//index:  0  cmd:  D:\software\JetBrains\GoWorkSpace\DailyTool\day02\01-switch.exe
	//index:  1  cmd:  hello
	//index:  2  cmd:  world
	//index:  3  cmd:  test
	//os.Args[0] ==> 程序名称
	//os.Args[1] ==> 第一个参数 一次类推
	cmds := os.Args
	for i, cmd := range cmds {
		fmt.Println("index: ", i, " cmd: ", cmd)
	}
	if len(cmds) < 2 {
		fmt.Println("请输入正确的参数！")
		return
	}

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		//go的switch 默认加了break了 不需要手工处理
		//如果想向下穿透的话 那么需要加上关键字: fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default called")
	}
}
