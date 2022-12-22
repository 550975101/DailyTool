package main

import (
	"fmt"
	"time"
)

func main() {
	//-单向读通道:
	//var numChanReadOnly <-chan int
	//-单向写通道
	//var numChanWriteOnly <-chan int

	//生产者消费者模型
	//Go: goroutine + channel

	//1、在主函数中创建一个双向通道 numChan
	numChan1 := make(chan int, 50)
	//双向通道可以赋值给同类型的单向通道 单向不能转双向
	//2、将numChan传递给producer  负责生产
	go producer(numChan1)
	//3、将numChan传递给consumer  负责消费
	go consumer(numChan1)

	time.Sleep(2 * time.Second)
	fmt.Println("OVER!")
}

// producer 生产者 ===》提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 60; i++ {
		out <- i
		//写通道不允许有读取操作
		//data := <-out
		fmt.Println("====>向管道中写入数据: ", i)
	}
}

// consumer 消费===》提供一个只读的通道
func consumer(in <-chan int) {
	//读通道不允许有写入操作
	//in <- 10
	for {
		v, ok := <-in
		if !ok {
			fmt.Println("通道已经关闭")
			break
		}
		fmt.Println("从管道中读取数据: ", v)
	}
}
