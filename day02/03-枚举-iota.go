package main

import "fmt"

// 在go语言中没有枚举类型 但是我们可以使用const+iota （常量累加器）来进行模拟
// 模拟一个一周的枚举
const (
	MONDAY    = iota //0
	TUESDAY   = iota //1
	WEDNESDAY        //2==》没有赋值 默认与上一行相同iota==》3
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
	M, N = iota, iota //7,7 const属于预编译期赋值的  所以不需要:=进行自动推导
)

const (
	JANU = iota + 1 //1
	FER             //2
	MAR             //3
	APRI            //4
)

/*
*
1、iota是常量组计数器
2、iota从0开始,没换行递增1
3、常量组有个特点如果不赋值，默认与上一行表达式相同
4、如果同一行出现两个iota，那么这两个iota的值是相同的
5、每个常量组的iota是独立的 如果遇到const iota会重新清零
*/
func main() {

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	fmt.Println(SATURDAY)
	fmt.Println(SUNDAY)
	fmt.Println(M)
	fmt.Println(N)
	fmt.Println(JANU)
	//var number int
	//var name string
	//var flag bool
	//可以使用变量组来将统一变量定义
	var (
		number int
		name   string
		flag   bool
	)
	fmt.Println("number: ", number, " name:", name, " flag：", flag)
}
