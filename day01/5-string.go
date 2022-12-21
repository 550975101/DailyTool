package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "duke"
	description := "汉字123qwer"
	//需要换行 原生输入字符串 使用反引号 ``
	usage := `你好：
    谢谢！`
	fmt.Println(name)
	fmt.Println(usage)

	//2、长度  访问
	//C++: name.length
	//GO:string没有.length方法 可以使用自由函数len()进行处理
	//len 很常用  统计的字符个数  一个汉字3个字符  统计汉字个数需要utf8.RuneCountInString
	l1 := len(name)
	fmt.Println("l1: ", l1)
	l3 := len(description)
	fmt.Println("l3: ", l3)
	l2 := utf8.RuneCountInString(description)
	fmt.Println("l2: ", l2)

	for i := 0; i < l3; i++ {
		fmt.Printf("i: %d,v: %c\n", i, description[i])
	}

	//3-拼接
	i, j := "hello", "world"
	fmt.Println("i+j=", i+j)

	//使用cons修饰的常量 不能修改
	const address = "beijing"
	//报错了 下面一行修改 编辑器就会提示
	//address = "shanghai"
	fmt.Println("address: ", address)
}
