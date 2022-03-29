package message

import "communicate_system-master/common/model"

// LoginMes 用户登陆发送消息
type LoginMes struct {
	UserID   int    `json:"userID"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

// RegisterMes 注册用户发送消息
type RegisterMes struct {
	User model.User `json:"user"`
}

// GroupMes 群聊消息
type GroupMes struct {
	Content    string `json:"content"` // 消息内容
	model.User        // 发送用户
}

// PrivateMes 私聊消息
type PrivateMes struct {
	Content  string     `json:"content"` // 消息内容
	Sender   model.User // 发送用户
	Receiver model.User // 接受用户
}

// OfflineMes 离线消息
type OfflineMes struct {
	CurUser model.User // 当前用户
}
