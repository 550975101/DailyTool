package main

import "fmt"

func main() {
	//1、定义一个字典
	//学生id ==》学生名字  id  name
	//定义一个map是不能直接赋值的  他是空的
	//panic: assignment to entry in nil map 对没有空间的map赋值 导致了异常
	//var idNames map[int]string
	//使用map之前 一定要手动分配空间
	//2、分配空间 make 指定10个长度，也可以不指定 建议直接指定长度 性能更好
	//idNames = make(map[int]string, 10)
	//一定建议定义直接分配空间 省的分两步
	idNames := make(map[int]string, 10)
	idNames[0] = "duke"
	idNames[1] = "lily"
	//4、遍历map
	for k, v := range idNames {
		fmt.Println("key: ", k, " value:", v)
	}
	//5、判断key是否包含
	//在map中不存在访问越界的问题 他认为所有的key都是有效的 所以访问一个不存在的key不会崩溃，返回这个类型的零值
	//零值: bool->false  数字->0 字符串->空
	name9 := idNames[9]
	fmt.Println("name9: ", name9)

	//通过这种方式来判断key是否存在
	value, ok := idNames[10]
	if ok {
		fmt.Println("id=10这个key存在的  value为", value)
	} else {
		fmt.Println("id=10这个key不存在  value为", value)
	}

	//6、删除map中元素
	//使用delete函数删除指定的key
	fmt.Println("idNames删除之前: ", idNames)
	delete(idNames, 1)
	fmt.Println("idNames删除之后: ", idNames)
	//删除一个存在的key 不会报错
	delete(idNames, 100)
	fmt.Println("idNames删除之后: ", idNames)
	//TODO
	//并发任务处理的时候 需要对map进行上锁
}
