package processor

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// ConnectWithServer 保持与服务端之间的通信
func ConnectWithServer(conn net.Conn) {
	// 不断读取服务端的消息
	msgTransfer := &utils.MsgTransfer{Conn: conn}
	for {
		mes, err := msgTransfer.ReadMsg()
		if err != nil {
			fmt.Println("transfer.ReadMsg err=", err)
			return
		}
		switch mes.Type {
		case constance.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes) // 更新并显示所有用户状态
		case constance.GroupMesType:
			var groupMes message.GroupMes
			json.Unmarshal([]byte(mes.Data), &groupMes)
			fmt.Printf("收到来自用户id为%d的群发消息,内容为%s\n", groupMes.UserID, groupMes.Content)
		case constance.PrivatMesType:
			var privateMes message.PrivateMes
			json.Unmarshal([]byte(mes.Data), &privateMes)
			fmt.Printf("收到来自用户id为%d的私聊消息,内容为%s\n", privateMes.Sender.UserID, privateMes.Content)
		case constance.OfflineResMesType:
			var offlineResMsg message.OfflineResMes
			json.Unmarshal([]byte(mes.Data), &offlineResMsg)
			onlineUsers[offlineResMsg.UserID].UserStatus = constance.UserOffline
		default:
			fmt.Println("服务端返回了未知的消息类型")
		}
	}
}

// ShowMenu 显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("--------------恭喜xxx登陆成功--------------")
	fmt.Println("\t 1.显示在线用户列表")
	fmt.Println("\t 2.发送群聊消息")
	fmt.Println("\t 3.发送私聊消息")
	fmt.Println("\t 4.退出系统")
	fmt.Println("-----------------------------------------")
	var key int
	var content string
	var userID int
	var sendMesProcess SendMesProcess
	fmt.Print("请选择1～4:")
	fmt.Scanf("%d", &key)
	fmt.Scanf("%d", &key)
	switch key {
	case 1:
		showOnlineUsers()
	case 2:
		fmt.Println("请输入你要发送的话:")
		fmt.Scanf("%s", &content)
		fmt.Scanf("%s", &content)
		sendMesProcess.SendGroupMes(content)
	case 3:
		fmt.Println("请输入你要私聊的用户id")
		fmt.Scanf("%d", &userID)
		fmt.Scanf("%d", &userID)
		fmt.Println("请输入你要发送的话:")
		fmt.Scanf("%s", &content)
		fmt.Scanf("%s", &content)
		sendMesProcess.SendPrivateMes(content, userID)
	case 4:
		fmt.Println("退出系统")
		sendMesProcess.SendOfflineMes()
		os.Exit(0)
	default:
		fmt.Println("您的选项不正确,请重新输入")
	}
}
