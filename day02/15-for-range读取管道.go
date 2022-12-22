package main

import (
	"fmt"
	"time"
)

/*
*
管道总结
1、当管道写满了 写阻塞
2、当缓冲区读完了 读阻塞
3、如果管道没有make分配空间 管道默认是nil
4、从nil的管道读取数据 写入数据 都会阻塞（注意不会奔溃）
5、从一个已经close的管道读取数据时，会返回零值（不会奔溃）
6、从一个已经close的管道写数据时，会奔溃
7、关闭一个已经close的管道 程序会奔溃
8、关闭管道的动作 一定要在写端 不应该放在读端
9、读和写的次数 一定是对等 否则:1、在多个go程中:资源泄露 2、在主go程中，程序奔溃(deadlock)
*/
func main() {
	numChan2 := make(chan int, 10)
	//写
	go func() {
		for i := 0; i < 50; i++ {
			numChan2 <- i
			fmt.Println("写入数据: ", i)
		}
		fmt.Println("数据全部写完,准备关闭管道")
		close(numChan2)
		//重复关闭 奔溃
		//close(numChan2)
		//关闭之后继续写 奔溃
		//numChan2 <- 10
	}()
	//遍历管道时,只返回一个值
	//for range是不知道管道是否已经写完 所以会一直在这里等待
	//在写入端 将管道关闭 for range遍历关闭的管道时 会退出
	for v := range numChan2 {
		fmt.Println("读取数据: ", v)
	}

	time.Sleep(3 * time.Second)
	//返回零值 不会奔溃
	i := <-numChan2
	fmt.Println("已经关闭之后 继续读取: ", i)

	fmt.Println("OVER!")
}
