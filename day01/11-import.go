package main

import (
	// "DailyTool/day01/sub" //sub是文件名 同时也是包名
	//SUB "DailyTool/day01/sub" //SUB是我们自己重新命名的包名
	. "DailyTool/day01/sub" //.代表用户在调用这个包里面的函数时，不需要使用包名.形式，不建议使用
	"fmt"
)

func main() {
	///包名.函数 去调用
	//s := sub.Sub(2, 2)
	//包名.函数 去调用 自己起的别名
	//s := SUB.Sub(2, 2)
	// . 直接函数名
	s := Sub(2, 2)
	fmt.Println("sub: ", s)
	//如果一个包里面的函数想要对外提供访问权限 那么一定要首写字母大写
	//大写字母开头的函数 相当于public
	//小写字母开头的函数 相当于private，只有相同包名的文件才能使用

}
