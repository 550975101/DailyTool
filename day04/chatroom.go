package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 将所有代码写在一个文件中 不做代码整理

type User struct {
	//名字
	name string
	//唯一的id
	id string
	//管道
	msg chan string
}

// 创建一个全局的map结构 用于保存所有的用户
var allUsers = make(map[string]User)

// 定义一个message全局通道 用于接收任何发送的消息
var message = make(chan string, 10)

func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listener: ", listener)
		return
	}
	//启动全局唯一的go程 负责监听message通道
	go broadcast()

	fmt.Println("服务器启动成功")
	for {
		fmt.Println("主go程监听中")
		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			return
		}
		//建立连接
		fmt.Println("建立连接成功!")
		//启动处理业务的go程
		go handler(conn)
	}
}

// 处理具体业务
func handler(conn net.Conn) {
	fmt.Println("启动业务...")
	//TODO 代码这里以后再具体实现 当前保留
	//客户端与服务器建立连接的时候 会有ip和port===》当成user的id
	cliendAddr := conn.RemoteAddr().String()
	fmt.Println("cliendAddr: ", cliendAddr)
	//创建user
	newUser := User{
		//id 我们不会修改,这个作为在map中的key
		id: cliendAddr,
		//可以修改 会提供rename命令修改 建立连接时 初始值与id相同
		name: cliendAddr,
		//注意需要make空间 否则无法写入数据
		msg: make(chan string, 10),
	}
	//添加用户到map结构
	allUsers[newUser.id] = newUser

	// 定义一个退出信号 用户监听client退出
	var isQuit = make(chan bool)
	//定义一个用于重置计数器的管道 用于告知watch函数 当前用户正在输入
	var restTimer = make(chan bool)
	//启动go程 负责监听退出信号
	go watch(&newUser, conn, isQuit, restTimer)

	//启动go程 负责将msg信息返回给客户端
	go writeBackToClient(&newUser, conn)

	//向message写入数据 当前用户上线的消息 用于通知所有人(广播)
	loginInfo := fmt.Sprintf("[%s]:[%s]==>上线了longin", newUser.id, newUser.name)

	message <- loginInfo
	for {
		buf := make([]byte, 1024)
		//读取客户端发送过来的请求数据
		cnt, err := conn.Read(buf)

		if cnt == 0 {
			fmt.Println("客户端主动关闭ctrl+c,准备退出")
			//map删除用户 conn  close掉
			//服务器可以主动的退出
			//在这里不进行真正的退出 而是发一个退出信号 统一做退出处理 可以使用新的管道来做信号传递
			isQuit <- true
		}
		if err != nil {
			fmt.Println(" conn read err: ", err)
			return
		}
		fmt.Println("服务器接收客户端发送过来的数据为: ", string(buf[:cnt]), " ,cnt", cnt)
		//-------------------------------------业务逻辑处理 开始-----------------------------------
		//1、查询当前所有的用户 who
		//这是用户输入的数据 最后一个是回车 我们去掉它
		userInput := string(buf[:cnt-1])
		//  a.先判断接收的数据是不是who ==> 长度&&字符串
		if len(userInput) == 3 && userInput == "who" {
			//  b.遍历allUser这个map:key:=userId value：user本身  将id和name拼接成一个字符串，返回给客户端
			fmt.Println("用户即将查询所有的用户信息")
			//这个切片包含了所有的用户信息
			var userInfos []string
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userid:%s,username:%s", user.id, user.name)
				userInfos = append(userInfos, userInfo)
			}
			r := strings.Join(userInfos, "\n")
			//将数据返回给查询的客户端
			newUser.msg <- r
		} else if len(userInput) > 9 && userInput[:7] == "rename" {
			//规则: rename|Duke
			//1、读取数据判断长度7 判断字符是rename
			//2、使用|进行分割 获取|后面的部分 作为名字
			array := strings.Split(userInput, "|")
			name := array[1]
			//3、更新用户名字newUser.name = Duke
			newUser.name = name
			//更新map中的user
			allUsers[newUser.id] = newUser
			//4、通知客户端 更新成功
			newUser.msg <- "rename successfully!"
		} else {
			//r如果用户输入的不是命令 只是普通地聊天信息 那么只需要写道广播通道即可 由其他的go程进行常规转发
			message <- userInput
		}
		restTimer <- true
		//-------------------------------------业务逻辑处理 结束----------------------------------
	}
}

func broadcast() {
	fmt.Println("广播go程启动成功...")
	defer fmt.Println("broadcast 程序退出!")
	for {
		//1、从message中读取数据
		fmt.Println("broadcast监听message中...")
		info := <-message
		fmt.Println("message 接收到消息: ", info)
		//2、将数据写入到每一个用户的msg管道中
		for _, user := range allUsers {
			//如果msg是非缓冲的，那么会在这里阻塞了
			user.msg <- info
		}
	}
}

// 每个用户应该还有一个用来监听自己msg通道的go程,负责将数据返回给客户端
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("user: %s 的go程正在监听自己的msg管道: ", user.name)
	for data := range user.msg {
		fmt.Printf("user: %s 写回给客户端数据为: %s \n", user.name, data)
		_, _ = conn.Write([]byte(data))
	}
}

// 启动一个go程 负责监听退出信号 触发后 进行清零的工作: delete map  close conn都在这里处理
func watch(user *User, conn net.Conn, isQuit <-chan bool, restTimer <-chan bool) {
	fmt.Println("启动监听退出信号的go程...")
	defer fmt.Println("watch go程退出!")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already\n", user.name)
			fmt.Println("删除当前用户")
			delete(allUsers, user.id)
			message <- logoutInfo
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout already!\n", user.name)
			fmt.Println("删除当前用户")
			delete(allUsers, user.id)
			message <- logoutInfo
			conn.Close()
			return
		case <-restTimer:
			fmt.Printf("连接%s 重置计数器!", user.name)
		}
	}
}
