package utils

import (
	"communicate_system-master/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// MsgTransfer 消息传送结构体
type MsgTransfer struct {
	Conn net.Conn   // 网络连接
	Buf  [8096]byte // 缓冲区
}

// ReadMsg 读取消息
func (m *MsgTransfer) ReadMsg() (mes message.Message, err error) {
	// 读取消息长度
	_, err = m.Conn.Read(m.Buf[:4])
	if err != nil {
		fmt.Println("conn.read(length) err=", err)
		return
	}
	length := binary.BigEndian.Uint32(m.Buf[0:4]) // 将byte[]转换为整数类型
	// 读取消息(根据消息长度)
	n, err := m.Conn.Read(m.Buf[:length])
	if uint32(n) != length || err != nil {
		fmt.Println("conn.read(msg) err=", err)
		return
	}
	// 反序列化消息
	err = json.Unmarshal(m.Buf[:length], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// SendMsg 发送消息(先发送消息长度、再发送消息本身)
func (m *MsgTransfer) SendMsg(mes *message.Message) (err error) {
	// 序列化消息
	msg, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 发送消息(先发送消息长度、再发送数据本身)
	var length [4]byte
	binary.BigEndian.PutUint32(length[0:4], uint32(len(msg))) // 将整数类型转换为byte[]
	n, err := m.Conn.Write(length[0:4])                       // 发送消息长度
	if n != 4 || err != nil {
		fmt.Println("conn.write(length) err=", err)
		return
	}
	_, err = m.Conn.Write(msg) // 发送消息
	if err != nil {
		fmt.Println("conn.write(msg) err=", err)
		return
	}
	return
}
