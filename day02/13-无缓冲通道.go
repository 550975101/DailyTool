package main

import (
	"fmt"
	"time"
)

func main() {
	//sync.RWMutex{}
	//当涉及多go程时  c语言使用上锁来保持资源同步  避免资源竞争问题
	//go语言也支持这种方式 但是go语言更好的解决方案时使用管道 通道
	//使用通道不需要我们去进行加锁
	//A往通道里写数据 B从管道里面读取数据 go自动帮我们做好了数据同步
	//创建管道 创建一个装有数字的管道 ===>channel
	//装数字管道一定要make  通map一样 否则是nil

	//无缓冲的管道
	//numChan := make(chan int)
	//有缓冲的管道
	numChan := make(chan int, 10)
	//儿子读数据
	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("data: ", data)
		}
	}()

	time.Sleep(5 * time.Second)
	//装字符串
	//strChan := make(chan string)
	//创建两个go程  父亲写数据 儿子读数据
	for i := 0; i < 50; i++ {
		numChan <- i
		fmt.Println("这是主go程，写数据: ", i)
	}

}
