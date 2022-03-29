package processor

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/common/model"
	"communicate_system-master/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// curUser 当前用户
var curUser CurUser

// CurUser 当前用户结构体
type CurUser struct {
	Conn net.Conn `json:"conn"`
	model.User
}

// UserProcess 用户处理结构体
type UserProcess struct {
}

// Login 用户登陆处理
func (u *UserProcess) Login(userID int, userPwd string) (err error) {
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接8889端口错误:err=", err)
		return
	}
	defer conn.Close()
	// 2.创建登陆消息
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd
	// 3.序列化登陆消息
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	// 4.创建消息
	var mes message.Message
	mes.Type = constance.LoginMesType
	mes.Data = string(data)
	// 5.发送消息
	msgTransfer := &utils.MsgTransfer{Conn: conn}
	err = msgTransfer.SendMsg(&mes)
	if err != nil {
		fmt.Println("sendMsg(conn) err=", err)
		return
	}
	// 6.读取服务端返回的消息
	msg, err := msgTransfer.ReadMsg()
	if err != nil {
		fmt.Println("readMsg(conn) err=", err)
		return
	}
	// 7.反序列化返回消息的data部分
	var loginResMsg message.LoginResMes
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	// 8.针对返回消息进行处理
	if loginResMsg.Code == 200 {
		fmt.Println("登陆成功")
		// 初始化curUser
		curUser.Conn = conn
		curUser.UserID = userID
		curUser.UserStatus = constance.UserOnline
		for _, id := range loginResMsg.UserIDs {
			if id == userID {
				continue
			}
			// 完成客户端onlineUsers初始化
			user := &model.User{
				UserID:     id,
				UserStatus: constance.UserOnline,
			}
			onlineUsers[id] = user
		}
		// 这里需要启动一个协程,用于保持与服务端之间的通信
		go ConnectWithServer(conn)
		// 循环显示登陆成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}
	return
}

// Register 用户注册处理
func (u *UserProcess) Register(userID int, userPwd string) (err error) {
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接8889端口错误:err=", err)
		return
	}
	defer conn.Close()
	// 2.创建注册消息
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	// 3.序列化注册消息
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	// 4.创建消息
	var mes message.Message
	mes.Type = constance.RegisterMesType
	mes.Data = string(data)
	// 5.发送消息
	msgTransfer := &utils.MsgTransfer{Conn: conn}
	err = msgTransfer.SendMsg(&mes)
	if err != nil {
		fmt.Println("注册消息发送错误err=", err)
	}
	// 6.读取服务端返回的消息
	msg, err := msgTransfer.ReadMsg()
	if err != nil {
		fmt.Println("ReadMsg(conn) err=", err)
		return
	}
	// 7.反序列化返回消息的data部分
	var registerResMsg message.RegisterResMes
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	// 8.针对返回消息进行处理
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功,请登陆")
	} else {
		fmt.Println(registerResMsg.Error)
		os.Exit(0)
	}
	return
}
