package processor

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/utils"
	"encoding/json"
	"fmt"
)

// SendMesProcess 发送消息处理器结构体
type SendMesProcess struct {
}

// ProcessGroupMes 处理群聊消息
func (s *SendMesProcess) ProcessGroupMes(mes *message.Message) {
	// 取出消息内容并反序列化
	var groupMes message.GroupMes
	err := json.Unmarshal([]byte(mes.Data), &groupMes)
	if err != nil {
		fmt.Println("json.Unmarshal err,", err)
		return
	}
	// 遍历服务器onlineUsers,将消息转发
	for id, userProcess := range onlineUserProcess.onlineUsers {
		if id == groupMes.UserID { // 过滤自己
			continue
		}
		// 发送消息
		msgTransfer := &utils.MsgTransfer{Conn: userProcess.Conn}
		err = msgTransfer.SendMsg(mes)
		if err != nil {
			fmt.Println("转发群聊消息失败err=", err)
		}
	}
}

// ProcessPrivateMes 处理私聊消息
func (s *SendMesProcess) ProcessPrivateMes(mes *message.Message) {
	// 取出消息内容并反序列化
	var privateMes message.PrivateMes
	err := json.Unmarshal([]byte(mes.Data), &privateMes)
	if err != nil {
		fmt.Println("json.Unmarshal err,", err)
		return
	}
	receiverID := privateMes.Receiver.UserID
	receiverProcessor, ok := onlineUserProcess.onlineUsers[receiverID]
	if !ok {
		fmt.Println("私聊用户不存在err=", err)
		return
	}
	msgTransfer := &utils.MsgTransfer{Conn: receiverProcessor.Conn}
	err = msgTransfer.SendMsg(mes)
	if err != nil {
		fmt.Println("发送消息失败err=", err)
	}
}

// ProcessOfflineMes 处理下线消息
func (s *SendMesProcess) ProcessOfflineMes(mes *message.Message) (err error) {
	// 取出消息内容并反序列化
	var offlineMes message.OfflineMes
	err = json.Unmarshal([]byte(mes.Data), &offlineMes)
	if err != nil {
		fmt.Println("json.Unmarshal err,", err)
		return
	}
	// 在线用户中删除该用户
	onlineUserProcess.DeleteOnlineUser(offlineMes.CurUser.UserID)
	// 包装返回消息
	var resMes message.Message
	resMes.Type = constance.OfflineResMesType
	var offlineResMes message.OfflineResMes
	offlineResMes.UserID = offlineMes.CurUser.UserID
	offlineResMes.UserStatus = constance.UserOffline
	data, err := json.Marshal(offlineResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)
	// 给其他用户发送返回消息
	for _, userProcess := range onlineUserProcess.onlineUsers {
		msgTransfer := &utils.MsgTransfer{Conn: userProcess.Conn}
		err = msgTransfer.SendMsg(&resMes)
		if err != nil {
			fmt.Println("notify me to other onlineUser err:", err)
			return
		}
	}
	return
}
