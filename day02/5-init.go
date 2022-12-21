package main

import _ "DailyTool/day02/sub"

func main() {
	//1、import "DailyTool/day02/sub"
	//sub.Sub(2, 1)
	//打印如下
	//this is first init() in package sub
	//this is second init() in package sub
	//this is test() in sub/utils ！

	//2、import _ "DailyTool/day02/sub"
	//打印如下
	//this is first init() in package sub
	//this is second init() in package sub
}
