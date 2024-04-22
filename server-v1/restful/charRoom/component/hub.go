package component

import "log"

type Hub struct {
	Clients    map[uint64]*Client // 客户端集合
	Broadcast  chan []byte        // 广播消息处理
	Register   chan *Client       // 注册chan
	Unregister chan uint64        // 注销chan
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[uint64]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan uint64),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// 注册客户端
			h.Clients[client.UserId] = client
		case userId := <-h.Unregister:
			if client, ok := h.Clients[userId]; ok {
				// 关闭客户端连接
				err := client.WsConn.Close()
				if err != nil {
					log.Println("关闭连接失败", err)
				}
				// 注销客户端
				delete(h.Clients, userId)
				// 关闭发送消息的channel
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for _, v := range h.Clients {
				v.Send <- message
			}
		}
	}
}
