package main

import "fmt"

func main() {
	p1 := testPtr()
	fmt.Println("*ptr", *p1)
	fmt.Println("ptr", p1)

	//如果一个包里面的函数要对外提供访问权限 函数名那么一定要首写字符大写: public
}

// 定义一个函数 返回string类型的指针 go语言返回写在参数列表后面
func testPtr() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
