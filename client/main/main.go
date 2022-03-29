package main

import (
	"communicate_system-master/client/processor"
	"fmt"
	"os"
)

func main() {
	// 接受用户的选择
	var key int
	var userID int
	var userPwd string

	for {
		fmt.Println("--------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t 1.登陆聊天室")
		fmt.Println("\t\t 2.注册用户")
		fmt.Println("\t\t 3.退出系统")
		fmt.Println("\t\t 请2选择 1～3")
		fmt.Println("---------------------------------------------")
		fmt.Scanf("%d", &key)
		fmt.Scanf("%d", &key)
		switch key {
		case 1:
			fmt.Print("请输入您的id:")
			fmt.Scanf("%d", &userID)
			fmt.Scanf("%d", &userID)
			fmt.Print("请输入您的password:")
			fmt.Scanf("%s", &userPwd)
			fmt.Scanf("%s", &userPwd)
			userProcessor := &processor.UserProcess{}
			userProcessor.Login(userID, userPwd)
		case 2:
			fmt.Print("请输入用户id:")
			fmt.Scanf("%d", &userID)
			fmt.Scanf("%d", &userID)
			fmt.Print("请输入用户password:")
			fmt.Scanf("%s", &userPwd)
			fmt.Scanf("%s", &userPwd)
			userProcessor := &processor.UserProcess{}
			userProcessor.Register(userID, userPwd)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误,请重新输入")
		}
	}
}
