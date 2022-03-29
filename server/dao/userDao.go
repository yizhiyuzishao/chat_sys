package dao

import (
	"communicate_system-master/common/model"
	"communicate_system-master/common/myError"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// MyUserDAO 全局userDAO对象
var MyUserDAO *UserDAO

// UserDAO 用户相关dao操作结构体
type UserDAO struct {
	Pool *redis.Pool // redis连接池
}

// NewUserDAO 创建userDAO实例
func NewUserDAO(pool *redis.Pool) (userDAO *UserDAO) {
	return &UserDAO{
		Pool: pool,
	}
}

// GetUserByID 根据id获取用户
func (u *UserDAO) GetUserByID(conn redis.Conn, id int) (user *model.User, err error) {
	res, err := redis.String(conn.Do("HGet", "user", id))
	if err != nil {
		if err == redis.ErrNil { // 表示在user哈希中,没有找到对应的id
			err = myError.ERROR_USER_NOT_EXISTS
		}
		return
	}
	user = &model.User{}
	err = json.Unmarshal([]byte(res), user) // 反序列化成user实例
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// InsertUser 新增用户到数据库
func (u *UserDAO) InsertUser(conn redis.Conn, id int, pwd string) (err error) {
	// 将用户信息序列化
	user := &model.User{UserID: id, UserPwd: pwd}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 将用户信息存入redis
	_, err = conn.Do("HSet", "user", user.UserID, string(data))
	if err != nil {
		fmt.Println("插入用户错误err=", err)
		return
	}
	return
}
