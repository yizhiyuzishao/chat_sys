package model

// User 用户结构体
type User struct {
	UserID     int    `json:"userID"`
	UserPwd    string `json:"userPwd"`
	UserStatus int    `json:"userStatus"` // 用户状态
}
