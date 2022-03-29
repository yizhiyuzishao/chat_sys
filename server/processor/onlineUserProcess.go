package processor

import "fmt"

// onlineUserProcess 全局在线用户处理器
var onlineUserProcess *OnlineUserProcess

// OnlineUserProcess 在线用户处理器结构体
type OnlineUserProcess struct {
	onlineUsers map[int]*UserProcess // key为用户id,value为用户处理器对象
}

// init 初始化onlineUserProcess
func init() {
	onlineUserProcess = &OnlineUserProcess{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUser 添加在线用户
func (o *OnlineUserProcess) AddOnlineUser(userProcess *UserProcess) {
	o.onlineUsers[userProcess.UserID] = userProcess
}

// DeleteOnlineUser 删除在线用户
func (o *OnlineUserProcess) DeleteOnlineUser(userID int) {
	delete(o.onlineUsers, userID)
}

// GetAllOnlineUser 获取当前所有在线用户
func (o *OnlineUserProcess) GetAllOnlineUser() map[int]*UserProcess {
	return o.onlineUsers
}

// GetOnlineUserByID 根据id返回在线用户的userProcess
func (o *OnlineUserProcess) GetOnlineUserByID(userID int) (userProcess *UserProcess, err error) {
	userProcess, ok := o.onlineUsers[userID]
	if !ok { // 要查找的用户当前不在线
		err = fmt.Errorf("用户%d不存在", userID)
		return
	}
	return
}
