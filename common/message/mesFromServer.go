package message

// LoginResMes 用户登录响应消息
type LoginResMes struct {
	Code    int    `json:"code"`
	Error   string `json:"myError"`
	UserIDs []int  `json:"userIDs"` // 在线用户id切片
}

// RegisterResMes 用户注册响应消息
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"myError"`
}

// OfflineResMes 用户离线响应消息
type OfflineResMes struct {
	UserID     int `json:"userID"`     // 用户id
	UserStatus int `json:"userStatus"` // 用户状态
}

// NotifyUserStatusMes 服务端推送用户状态的消息
type NotifyUserStatusMes struct {
	UserID     int `json:"userID"`     // 用户id
	UserStatus int `json:"userStatus"` // 用户状态
}
