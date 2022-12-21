package main

import (
	"fmt"
)

func main() {
	//1-定义 定义一个具有是个数字的数组
	//c语言定义: int nums[10] ={1,2,3}
	//go语言定义: nums := [10]int{1,2,3,4,5} 常用方式
	//var nums =[10]int{1,2,3,4,5,6}
	//var nums [10]int =[10]int{1,2,3,4,5,6}
	nums := [10]int{1, 2, 3, 4}
	//2-遍历 方式一
	for i := 0; i < len(nums); i++ {
		fmt.Println("i：", i, "，j:", nums[i])
	}
	//方式二: for range ===>python支持
	for i, v := range nums {
		fmt.Println("index: ", i, "，value: ", v)
	}
	//如果需要忽略index 则
	//for _, v := range nums {
	//	fmt.Println( "value: ", v)
	//}

}
