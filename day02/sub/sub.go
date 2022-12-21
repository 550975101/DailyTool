package sub

import "fmt"

// 1、init函数没有参数 没有返回值 原型固定如下
// 2、一个包中包含多个init时  调用顺序是不确定的
// 3、init函数不允许用户显式调用
// 4、有的时候引用一个包 可能只是想使用这个包里面的init函数(mysql的init对驱动初始化)
// 但是不想使用这个包里面的其他函数 为了防止编译器报错 可以使用_形式来处理
func init() {
	fmt.Println("this is first init() in package sub")
}

func init() {
	fmt.Println("this is second init() in package sub")
}

func Sub(a, b int) int {
	test()
	return a - b
}
