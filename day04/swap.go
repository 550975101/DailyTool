package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Job string
}

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

func swap(pa *int, pb *int) {
	*pa, *pb = *pb, *pa
}

func main() {
	student := Student{
		Person: Person{},
		Job:    "",
	}
	fmt.Println(student.Age)
	a := 10
	b := 20

	fmt.Println("a = ", a, "b = ", b)
	// swap
	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b)

	params := make(map[string]interface{}, 10)
	params["key"] = "values"
	fmt.Println(params)

	//包含汉字字符串的遍历
	str := "hello北京"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符串=%c\n", r[i])
	}
	//字符串转整数
	num, err := strconv.Atoi("11")
	if err != nil {
		fmt.Println(" strconv.Atoi：", err)
		return
	}
	fmt.Println(num)
	//整数转字符串
	str2 := strconv.Itoa(12345)
	//打印类型%T
	fmt.Printf("%T", str2)

	//字符串转字节数组
	bytes := []byte("hello")
	fmt.Printf("bytes=%v\n", bytes)

	//字节数组转字符串
	str1 := string([]byte{97, 98, 99})
	fmt.Println(str1)

	//10进制转2，8，16进制
	num1 := strconv.FormatInt(123, 16)
	fmt.Println(num1)

	//查找子串是否在指定的字符串中
	contains := strings.Contains("seafood", "food")
	fmt.Println(contains)

	//统计字符串中有几个指定的子串
	count := strings.Count("cehese", "e")
	fmt.Println(count)

	//不区分大小写的字符串比较
	fold := strings.EqualFold("abc", "ABC")
	fmt.Println(fold)
	fmt.Println("abc" == "ABC")

	//返回子串在字符中第一个出现的index值 如果没有返回-1
	indexRune := strings.Index("中国人", "人")
	fmt.Println(indexRune)

	builder := &strings.Builder{}
	for i := 0; i < 10; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	fmt.Println(builder.String())

	fi, err := os.Open("D:\\software\\JetBrains\\GoWorkSpace\\DailyTool\\day04\\chatroom.go")
	if err != nil {
		fmt.Println(err)
	}
	defer func(fi *os.File) {
		err := fi.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fi)
	//创建Reader
	reader := bufio.NewReader(fi)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Format("20060102 15:04:05"))

	parse, err := time.Parse("20060102 15:04:05", "20221229 10:30:30")
	if err != nil {
		return
	}
	fmt.Println(parse.String())

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err: ", err)
		}
	}()

	strArr := [...]string{"1", "2", "3"}
	fmt.Printf("%T", strArr)
	for _, s := range strArr {
		fmt.Println(s)
	}
}

// 函数去读取配置文件init.conf
// 如果文件名传入不正确 我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
