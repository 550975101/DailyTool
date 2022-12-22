package main

import (
	"fmt"
	"time"
)

// 当程序中有多个channel协同工作 ch1 ch2，某一个时刻 ch1或ch2出发了 程序要做出相应的处理
// 使用select来监听多个通道 当管道触发时 (写入数据 读取数据 关闭管道)
// select语法与switch case很像 但是所有的分支条件必须是通道io
func main() {
	//var ch1 ch2  chan int
	ch1, ch2 := make(chan int, 10), make(chan int, 10)

	//启动一个go程 负责监听两个channel
	go func() {
		for {
			fmt.Println("监听中......")
			select {
			case data1 := <-ch1:
				fmt.Println("从ch1读取数据成功，data1: ", data1)
			case data2 := <-ch2:
				fmt.Println("从ch2读取数据成功，data2: ", data2)
			default:
				fmt.Println("select default分支called")
				time.Sleep(time.Second)
			}
		}
	}()
	//启动go1 写ch1
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second / 2)
			fmt.Println("ch1 写入数据: ", i)
		}
	}()
	//启动go 写ch2
	go func() {
		for i := 11; i < 20; i++ {
			ch2 <- i
			time.Sleep(1 * time.Second / 2)
			fmt.Println("ch2 写入数据: ", i)
		}
	}()

	for {

	}
}
