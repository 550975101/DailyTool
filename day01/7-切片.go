package main

import "fmt"

func main() {
	//切片: slice,他的底层也是数组 可以动态改变长度
	//定义一个切片 包含多个地名
	names := []string{"北京", "上海", "广州", "深圳", "洛阳", "南京", "廊坊"}
	for i, v := range names {
		fmt.Println("i: ", i, " v: ", v)
	}

	//1、追加数据
	names1 := append(names, "海南")
	fmt.Println("names: ", names)
	fmt.Println("names1: ", names1)

	fmt.Println("追加元素前，name的长度: ", len(names), "容量: ", cap(names))
	names = append(names, "海南")
	fmt.Println("追加元素后，name的长度: ", len(names), "容量: ", cap(names))
	fmt.Println("names追加元素后赋值给自己: ", names)
	fmt.Println("names1: ", names1)

	//2、对于一个切片 不仅是有长度的概念len()  还有一个容量的概念cap()
	var nums []int
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("len: ", len(nums), " 容量: ", cap(nums))
	}
	//想基于names创建一个新的数组
	names2 := [3]string{}
	names2[0] = names[0]
	names2[1] = names[1]
	names2[2] = names[2]

	//切片可以基于一个数组 灵活地创建新的数组
	names3 := names[0:3]
	fmt.Println("names3: ", names3)
	names3[0] = "Hello"
	fmt.Println("names3: ", names3)
	fmt.Println("names: ", names)

	//如果想让切片完全独立于原数组 可以使用copy()函数来完成

	//1、如果从第0个元素开始截取,那么冒号左边的数字可以省略
	//0-4 不包括索引4   左闭右开
	names4 := names[:4]
	fmt.Println("names4: ", names4)
	//1-最后 索引1开始到结束
	names5 := names[1:]
	fmt.Println("names5: ", names5)
	//截取全部
	names6 := names[:]
	fmt.Println("names6: ", names6)

	//4、也可以一基于一个字符串进行切片截取  取得字符串: helloWorld
	sub1 := "helloWorld"[5:7]
	fmt.Println("sub1: ", sub1)

	//5、可以在创建空切片的时候 明确指定切片的数量  这样可以提高运行效率
	//创一个长度是10  容量是20的string类型切片
	str2 := make([]string, 10, 20)
	fmt.Println("str2: len", len(str2), " cap:", cap(str2))
	str2[0] = "hello"
	str2[1] = "world"

	//6、如果想让切片完成独立于原始数组 可以使用copy()函数来完成
	namesCopy := make([]string, len(names))
	//func copy(dst,src []Type) int
	//names是一个数组 copy函数接收的类型是切片 所以使用[:]将数组变成切片
	copy(namesCopy, names)
	namesCopy[0] = "香港"
	fmt.Println("namesCopy: ", namesCopy)
	fmt.Println("names: ", names)
}
