package main

import "fmt"

type Human struct {
	//成员属性
	name   string
	age    int
	gender string
	score  float64
}

// Eat 在外面绑定方法
func (h *Human) Eat() {
	fmt.Println("this is: ", h.name)
}

type Student1 struct {
	hum Human //包含Human类型的变量
	//学校
	school string
}

// 定义一个老师继承Human
type Teacher struct {
	Human //直接写Human类型 没有字段名字
	//学科
	subject string
}

func main() {
	s1 := Student1{
		hum: Human{
			name:   "Lily",
			age:    18,
			gender: "女生",
		},
		school: "衡水中学",
	}
	fmt.Println("s1.name ", s1.hum.name)
	fmt.Println("s1.school ", s1.school)

	t1 := Teacher{}
	t1.subject = "语文"
	//下面的这几个字段是继承紫Human的
	t1.name = "李老师"
	t1.age = 10
	t1.gender = "女生"
	fmt.Println("t1: ", t1)
	t1.Eat()

	//继承的时候 虽然我们没有定义字段名字 但是会自动创建一个默认字段的同名字段
	//这是为了子类中依然可以操作父类 因为: 子类可能出现同名的字段
	fmt.Println("t1.Human.name: ", t1.Human.name)
}
