# communicate_system

#### 介绍
>基于C/S架构的采用纯Go语言+Redis实现的即时聊天系统

 **实现功能** ：
+ 用户注册
+ 用户登陆
+ 显示在线用户列表(如果一个登录用户离线，就把这个人从在线列表去掉)
+ 群聊
+ 私聊

 **运行方法** ：
+ 首先运行`service/main/main.go`启动服务端
+ 然后运行`client/main/main.go`启动客户端(可启动多个客户端)


