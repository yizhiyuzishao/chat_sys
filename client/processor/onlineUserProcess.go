package processor

import (
	"communicate_system-master/common/constance"
	"communicate_system-master/common/message"
	"communicate_system-master/common/model"
	"fmt"
)

// onlineUsers 维护在线用户(客户端)
var onlineUsers = make(map[int]*model.User, 10)

// updateUserStatus 更新并显示用户状态
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok { // 如果在线用户列表中没有该用户
		user = &model.User{
			UserID: notifyUserStatusMes.UserID,
		}
	}
	user.UserStatus = notifyUserStatusMes.UserStatus
	onlineUsers[notifyUserStatusMes.UserID] = user
	showOnlineUsers()
}

// showOnlineUsers 客户端显示当前在线用户
func showOnlineUsers() {
	fmt.Println("当前在线用户列表为:")
	for id, _ := range onlineUsers {
		if onlineUsers[id].UserStatus == constance.UserOnline {
			fmt.Println("用户id:\t", id)
		}
	}
}
