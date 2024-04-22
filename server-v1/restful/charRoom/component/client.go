package component

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	Hub    *Hub            // 客户端所属的hub
	UserId uint64          // 客户端id
	WsConn *websocket.Conn // websocket连接
	Send   chan []byte     // 发送消息的channel
}

func NewClient(hub *Hub, userId uint64, wsConn *websocket.Conn) *Client {
	return &Client{
		Hub:    hub,
		UserId: userId,
		WsConn: wsConn,
		Send:   make(chan []byte, 256),
	}
}

func (client *Client) WritePump() {
	// 开启一个定时器，每隔一段时间向客户端发送一个ping消息
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		err := client.WsConn.Close()
		if err != nil {
			log.Println("关闭连接失败", err)
		}
		log.Printf("客户端主动关闭连接: %d", client.UserId)
	}()

	// 监听channel
	for {
		select {
		// 监听发送消息的channel：Send（有消息就推送给客户端）
		case message, ok := <-client.Send:
			// 判断通道是否关闭
			if !ok {
				// 如果通道关闭，发送关闭消息，通知客户端关闭连接
				client.WsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 设置写超时时间
			client.WsConn.SetWriteDeadline(time.Now().Add(writeWait))

			// 写消息
			client.WsConn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			// 设置写超时时间
			client.WsConn.SetWriteDeadline(time.Now().Add(writeWait))
			// 发送ping消息
			if err := client.WsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) ReadPump() {
	// 结束时关闭连接，注销客户端
	defer func() {
		client.Hub.Unregister <- client.UserId
	}()

	client.WsConn.SetReadLimit(maxMessageSize) // 设置读取消息的最大长度

	// ping每隔9s发送一次，如果10s内没有收到pong消息，就关闭连接
	client.WsConn.SetReadDeadline(time.Now().Add(pongWait)) // 设置读消息的超时时间

	// 设置pong消息处理函数
	client.WsConn.SetPongHandler(func(string) error {
		// 如果收到pong消息，重置读消息的超时时间
		client.WsConn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		var message []byte
		_, message, err := client.WsConn.ReadMessage() // 读取消息

		// 判断是否读取消息失败
		if err != nil {
			// 如果读取消息失败，判断是否是连接被关闭的错误
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			// 中止循环
			break
		}

		// 将json消息解析为Message结构体
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("json unmarshal error:", err)
			return
		}

		// 根据消息类型进行不同的处理
		switch msg.Type {
		case 1:
			if v, ok := client.Hub.Clients[msg.Receiver]; ok {
				// 如果接收者在线，发送消息
				v.Send <- message
				// 给发送者返回发送成功的消息
				client.Send <- SuccessSend(msg.Sender, msg.Receiver)
			} else {
				// 如果接收者不在线，给发送者返回不在线的消息
				client.Send <- NotOnline(msg.Sender, msg.Receiver)
			}
		default:
			client.Hub.Broadcast <- message
		}
	}
}
