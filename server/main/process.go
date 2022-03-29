package main

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/server/processor"
	"communicate_system-master/utils"
	"fmt"
	"io"
	"net"
)

// Processor 总处理器结构体
type Processor struct {
	Conn net.Conn // 网络连接
}

// process 处理客户端与服务端之间的消息通信
func (p *Processor) process() {
	defer p.Conn.Close()
	// 循环读取客户端发送的消息
	for {
		msgTransfer := &utils.MsgTransfer{
			Conn: p.Conn,
		}
		msg, err := msgTransfer.ReadMsg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("数据读取结束...")
				return
			} else {
				fmt.Println("read data from client err=", err)
			}
		}
		// 处理读取到的消息
		err = p.serverProcessMes(&msg)
		if err != nil {
			return
		}
	}
}

// serverProcessMes	消息处理总控函数
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	// 处理登陆消息
	case constance.LoginMesType:
		userProcess := &processor.UserProcess{Conn: p.Conn}
		userProcess.ServerProcessLogin(mes)
	// 处理注册消息
	case constance.RegisterMesType:
		userProcess := &processor.UserProcess{Conn: p.Conn}
		userProcess.ServerProcessRegister(mes)
	// 处理群聊消息
	case constance.GroupMesType:
		sendMesProcess := &processor.SendMesProcess{}
		sendMesProcess.ProcessGroupMes(mes)
	// 处理私聊消息
	case constance.PrivatMesType:
		sendMesProcess := &processor.SendMesProcess{}
		sendMesProcess.ProcessPrivateMes(mes)
	// 处理退出消息
	case constance.OfflineMesType:
		sendMesProcess := &processor.SendMesProcess{}
		sendMesProcess.ProcessOfflineMes(mes)
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}
