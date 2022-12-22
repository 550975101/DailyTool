package main

import (
	"fmt"
	"time"
)

func main() {
	//numsChan := make(chan int, 10)
	//1、当缓冲写满的时候 写阻塞  当读取后 再恢复写入
	//2、当缓冲区读取完毕 读阻塞
	//3、如果管道没有使用make分配空间 那么管道默认是nil的 读取、写入都会阻塞
	//4、对于一个管道 读与写的次数 必须对等
	var names chan string
	names = make(chan string, 10)

	names <- "hello"
	fmt.Println("names: ", <-names)

	//1、当管道读写次数不一致的时候
	//如果阻塞在主go程 那么程序会奔溃
	//如果阻塞在子go程 那么程序会出现内存泄露

	numsChan := make(chan int, 10)
	//写
	go func() {
		for i := 0; i < 70; i++ {
			numsChan <- i
			fmt.Println("写入数据: ", i)
		}
	}()
	//读
	go func() {
		for i := 0; i < 60; i++ {
			fmt.Println("主程序准备读取数据.........")
			data := <-numsChan
			fmt.Println("读取数据: ", data)
		}
	}()

	for {
		fmt.Println("这是主go程，正在死循环")
		time.Sleep(1 * time.Second)
	}
}
