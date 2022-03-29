package processor

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/common/myError"
	"communicate_system-master/server/service"
	"communicate_system-master/utils"
	"encoding/json"
	"fmt"
	"net"
)

// UserProcess 用户处理结构体
type UserProcess struct {
	Conn   net.Conn // 网络连接
	UserID int      // 用户id
}

// ServerProcessLogin 处理登录消息
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 反序列化登录消息
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	// 用户信息校验,返回一个登录响应消息
	var loginResMsg message.LoginResMes
	user, err := service.LoginOA(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == myError.ERROR_USER_NOT_EXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == myError.ERROR_USER_PWD {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器内部错误"
		}
	} else {
		loginResMsg.Code = 200
		loginResMsg.Error = ""
		// 登陆成功,将用户放到onlineUsers中
		u.UserID = loginMes.UserID
		onlineUserProcess.AddOnlineUser(u)
		// loginResMsg中记录所有在线用户id
		for id, _ := range onlineUserProcess.onlineUsers {
			loginResMsg.UserIDs = append(loginResMsg.UserIDs, id)
		}
		// 推送其他在线用户我上线
		u.NotifyMeToOtherOnlineUser(loginMes.UserID)
		fmt.Println(user, "登陆成功")
	}
	msg, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 包装返回消息
	var resMes message.Message
	resMes.Type = constance.LoginResMesType
	resMes.Data = string(msg)
	// 发送返回消息
	msgTransfer := &utils.MsgTransfer{
		Conn: u.Conn,
	}
	msgTransfer.SendMsg(&resMes)
	return
}

// ServerProcessRegister 处理注册消息
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 反序列化注册消息
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	// 注册用户信息到redis,返回一个注册响应消息
	var registerResMes message.RegisterResMes
	err = service.Register(registerMes.User.UserID, registerMes.User.UserPwd)
	if err != nil {
		if err == myError.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
		registerResMes.Error = "用户注册成功"
	}
	msg, err := json.Marshal(registerResMes)
	if err != nil {
		if err != nil {
			fmt.Println("json.Marshal err=", err)
			return
		}
	}
	// 包装响应消息
	var resMes message.Message
	resMes.Type = constance.RegisterResMesType
	resMes.Data = string(msg)
	// 发送响应消息
	msgTransfer := &utils.MsgTransfer{
		Conn: u.Conn,
	}
	msgTransfer.SendMsg(&resMes)
	return
}

// NotifyMeToOtherOnlineUser 通知其他所有在线用户我上线了
func (u *UserProcess) NotifyMeToOtherOnlineUser(userID int) {
	// 遍历onlineUsers,逐个发送
	for id, userProcess := range onlineUserProcess.GetAllOnlineUser() {
		if id == userID { // 过滤自己
			continue
		}
		// 包装推送消息
		var mes message.Message
		mes.Type = constance.NotifyUserStatusMesType
		var notifyUserStatusMes message.NotifyUserStatusMes
		notifyUserStatusMes.UserID = userID
		notifyUserStatusMes.UserStatus = constance.UserOnline
		data, err := json.Marshal(notifyUserStatusMes)
		if err != nil {
			fmt.Println("json.Marshal err=", err)
			return
		}
		mes.Data = string(data)
		// 发送消息
		msgTransfer := &utils.MsgTransfer{Conn: userProcess.Conn}
		err = msgTransfer.SendMsg(&mes)
		if err != nil {
			fmt.Println("notify me to other onlineUser err:", err)
			return
		}
	}
}
