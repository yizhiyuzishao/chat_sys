package service

import (
	"communicate_system-master/common/model"
	"communicate_system-master/common/myError"
	"communicate_system-master/server/dao"
)

// LoginOA 登陆校验
func LoginOA(userID int, userPwd string) (user *model.User, err error) {
	conn := dao.MyUserDAO.Pool.Get()
	defer conn.Close()
	user, err = dao.MyUserDAO.GetUserByID(conn, userID)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = myError.ERROR_USER_PWD // 用户密码错误
		return
	}
	return
}

// Register 注册处理
func Register(userID int, userPwd string) (err error) {
	conn := dao.MyUserDAO.Pool.Get()
	defer conn.Close()
	// 判断用户是否存在
	_, err = dao.MyUserDAO.GetUserByID(conn, userID)
	if err == nil {
		err = myError.ERROR_USER_EXISTS // 用户不存在错误
		return
	}
	// 将用户信息存入redis
	err = dao.MyUserDAO.InsertUser(conn, userID, userPwd)
	if err != nil {
		return
	}
	return
}
