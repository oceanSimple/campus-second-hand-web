package component

import "encoding/json"

const (
	CONNECT       = iota // 建立连接
	SEND                 // 1 to 1 发送消息
	SystemMessage        // 系统消息
	BROADCAST            // 广播消息
	CLOSE                // 关闭连接
)

const (
	SuccessSendMsg = "success send"
	NotOnlineMsg   = "not online"
)

type Message struct {
	Type     uint64 `json:"type"`     // 消息类型
	Sender   uint64 `json:"sender"`   // 发送者的id
	Receiver uint64 `json:"receiver"` // 接收者的id
	Content  string `json:"content"`  // 消息内容
}

func SuccessSend(sender, receiver uint64) []byte {
	var msg = Message{
		Type:     SystemMessage,
		Content:  SuccessSendMsg,
		Sender:   sender,
		Receiver: receiver,
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return bytes
}

func NotOnline(sender, receiver uint64) []byte {
	var msg = Message{
		Type:     SystemMessage,
		Content:  NotOnlineMsg,
		Sender:   sender,
		Receiver: receiver,
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return bytes
}
