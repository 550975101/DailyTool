package main

import "fmt"

// 函数返回值在参数列表之后 如果有多个返回值 需要使用括号()包裹,多个参数之间使用,分割
func test(a int, b int, c string) (int, string) {
	return a + b, c
}

func test1(a, b int, c string) (code int, msg string) {
	//直接使用返回值的变量名参与运算
	code = a + b
	msg = c
	//当返回值有名字的时候，可以直接简写return
	//return code, msg
	return
}

// 当返回值只有一个的时候 并且没有名字 返回值位置不需要加圆括号
func test2(a, b int, c string) int {
	return 10
}

func main() {
	v1, s1 := test(10, 20, "hello")
	fmt.Println("v1: ", v1, " s1: ", s1)

	v2, s2 := test1(10, 20, "hello")
	fmt.Println("v2: ", v2, " s2: ", s2)
}
