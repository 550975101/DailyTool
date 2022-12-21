package main

import "fmt"

//Person类 绑定方法 Eat  Run Laugh 成员
//public private

type Person struct {
	//成员属性
	name   string
	age    int
	gender string
	score  float64
}

// Eat 在类外面绑定方法
// 使用非指针的话 原来的值不会改变
func (p Person) Eat() {
	fmt.Println("Person is eating")
	//类的方法 可以使用自己的成员
	fmt.Println(p.name + "is eating")
	p.name = "Duke"
}

// Eat2 注意 指针
// 使用指针的话 原来的值就会改变
func (p *Person) Eat2() {
	fmt.Println("Person is eating")
	//类的方法 可以使用自己的成员
	p.name = "Duke"
}

func main() {
	lily := Person{
		name:   "Lily",
		age:    30,
		gender: "女生",
		score:  10,
	}
	fmt.Println("lily: ", lily)
	//调用eat
	lily.Eat()
	//调用eat2
	lily.Eat2()
	fmt.Println("lily: ", lily)
}
