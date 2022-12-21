package main

import "fmt"

// 实现go多态 需要实现定义的接口
// 人类的武器发起攻击 不同等级的子弹效果不同
// 定义一个接口 注意类型是interface

// IAttack 1、定义一个接口 里面有设计好需要的接口 可以有多个
// 2、任何实现了这个接口的类型 都可以赋值给这个接口 从而实现多态
// 3、多个类之间不需要继承关系
// 4、如果interface中定义了多个接口 那么实际的类必须全部实现接口函数  才可以赋值
type IAttack interface {
	//接口函数可以有多个 但是只能有函数原型 不可以有实现
	Attack()
	//Attack1()
}

// HumanLowLevel 低等级
type HumanLowLevel struct {
	name  string
	level int
}

// HumanHighLevel 低等级
type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为: ", a.level, "造成1000点伤害")
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为: ", a.level, "造成5000点伤害")
}

// DoAttack 定义一个多态的通用接口 传入不同的对象  实现不同的效果
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {

	//var player interface{}
	//定义一个包含attack的接口变量
	var player IAttack

	lowLevel := HumanLowLevel{
		name:  "David",
		level: 1,
	}

	hightLevel := HumanHighLevel{
		name:  "David",
		level: 10,
	}

	lowLevel.Attack()
	//对player进行赋值为lowLevel  接口需要使用指针类型来赋值
	player = &lowLevel
	player.Attack()

	fmt.Println("多态..........")
	DoAttack(&lowLevel)
	DoAttack(&hightLevel)
}
