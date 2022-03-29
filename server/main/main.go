package main

import (
	"communicate_system-master/server/dao"
	"time"
	"fmt"
	"net"
)

func main() {
	// 初始化redis连接池
	initRedisPool("localhost:6379", 16, 0, 300*time.Second)
	// 新建userDAO对象
	dao.MyUserDAO = dao.NewUserDAO(pool)
	// 启动服务器监听8889端口
	fmt.Println("服务端在8889端口监听......")
	listener, err := net.Listen("tcp", "localhost:8889")
	defer listener.Close()
	// 监听失败
	if err != nil {
		fmt.Println("服务器监听错误err=", err)
		return
	}
	// 监听成功
	for {
		fmt.Println("等待客户端连接服务器.......")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端连接错误err=", err)
		}
		// 创建一个消息总控
		p := &Processor{Conn: conn}
		go p.process()
	}
}
