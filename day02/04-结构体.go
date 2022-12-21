package main

import "fmt"

type MyInt int //type相当于typef

// Person go语言结构体使用type+struct来处理
type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt
	i, j = 10, 20
	fmt.Println("i: ", i, " j:", j)

	//使用结构体各个字段
	lily := Student{
		name:   "Lily",
		age:    0,
		gender: "女生",
		score:  80, //最后一个元素后面必须加上逗号 如果不加逗号则必须与}同一行
	}
	fmt.Println("lily: ", lily.name, lily.age, lily.gender, lily.score)

	s1 := &lily
	fmt.Println("lily 使用指针s1.name打印: ", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("lily 使用指针(*s1).name打印: ", (*s1).name, s1.age, s1.gender, s1.score)

	//在定义期间对结构体赋值时,如果每个字段都赋值了 那么字段的名字可以省略不写
	Duke := Student{
		"Duke",
		28,
		"男生",
		99,
	}
	fmt.Println("Duke: ", Duke.name, Duke.age)
}
