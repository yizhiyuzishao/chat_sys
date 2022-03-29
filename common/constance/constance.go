package constance

// 消息类型常量
const (
	LoginMesType    = "LoginMes"    // 登陆发送消息
	LoginResMesType = "LoginResMes" // 登陆返回消息

	RegisterMesType    = "RegisterMes"    // 注册发送消息
	RegisterResMesType = "RegisterResMes" // 注册返回消息

	OfflineMesType    = "OfflineMes"    // 离线消息
	OfflineResMesType = "OfflineResMes" // 离线返回消息

	NotifyUserStatusMesType = "NotifyUserStatusMes" // 服务器推送用户状态消息

	GroupMesType  = "GroupMes"   // 群聊消息
	PrivatMesType = "PrivateMes" // 私聊消息

)

// 用户状态常量
const (
	UserOnline     = iota // 在线
	UserOffline           // 离线
	UserBusyStatus        // 忙线
)
