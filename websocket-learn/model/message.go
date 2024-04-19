package model

const (
	CONNECT = iota // 建立连接
	SEND           // 发送消息
	RECEIVE        // 接收消息
	CLOSE          // 关闭连接
)

type Message struct {
	Type     uint64 `json:"type"`     // 消息类型
	Sender   uint64 `json:"sender"`   // 发送者的id
	Receiver uint64 `json:"receiver"` // 接收者的id
	Content  string `json:"content"`  // 消息内容
}
