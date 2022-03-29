package message

// Message 通用消息结构体定义
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}
