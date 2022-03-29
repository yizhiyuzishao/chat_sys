package processor

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/utils"
	"encoding/json"
	"fmt"
)

// SendMesProcess 发送消息处理器
type SendMesProcess struct {
}

// SendGroupMes 发送群聊消息
func (s *SendMesProcess) SendGroupMes(content string) (err error) {
	// 包装发送消息
	var mes message.Message
	mes.Type = constance.GroupMesType
	var sendMes message.GroupMes
	sendMes.Content = content
	sendMes.UserID = curUser.UserID
	sendMes.UserStatus = curUser.UserStatus
	data, err := json.Marshal(sendMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	// 发送群聊消息给服务器
	msgTransfer := &utils.MsgTransfer{Conn: curUser.Conn}
	err = msgTransfer.SendMsg(&mes)
	if err != nil {
		fmt.Println("send message err=", err)
		return
	}
	return
}

// SendPrivateMes 发送私聊消息
func (s *SendMesProcess) SendPrivateMes(content string, userID int) (err error) {
	// 包装发送消息
	var mes message.Message
	mes.Type = constance.PrivatMesType
	var sendMes message.PrivateMes
	sendMes.Content = content
	sendMes.Sender.UserID = curUser.UserID
	sendMes.Sender.UserStatus = curUser.UserStatus
	sendMes.Receiver.UserID = userID
	data, err := json.Marshal(sendMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	// 发送群聊消息给服务器
	msgTransfer := &utils.MsgTransfer{Conn: curUser.Conn}
	err = msgTransfer.SendMsg(&mes)
	if err != nil {
		fmt.Println("send message err=", err)
		return
	}
	return
}

// SendOfflineMes 发送离线消息
func (s *SendMesProcess) SendOfflineMes() (err error) {
	// 包装发送消息
	var mes message.Message
	mes.Type = constance.OfflineMesType
	var offlineMes message.OfflineMes
	offlineMes.CurUser = curUser.User
	data, err := json.Marshal(offlineMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	// 发送消息
	msgTransfer := &utils.MsgTransfer{Conn: curUser.Conn}
	err = msgTransfer.SendMsg(&mes)
	if err != nil {
		fmt.Println("send message err=", err)
		return
	}
	return
}
