package model

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type Connection struct {
	WsConn *websocket.Conn // websocket连接
	Sc     chan []byte     // 发送消息的channel
	UserId uint64          // 用户id
}

// 持续监听管道消息，并将消息写入websocket连接
func (c *Connection) writer() {
	for message := range c.Sc {
		err := c.WsConn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	// 当连接的sc管道关闭时，关闭websocket连接
	err := c.WsConn.Close()
	if err != nil {
		panic(err)
	}
}

// 持续监听websocket连接消息，并将消息写入管道
func (c *Connection) reader() {
	for {
		// 第一个参数为消息类型(text or binary)，第二个参数为消息内容
		_, msg, err := c.WsConn.ReadMessage()
		if err != nil {
			// 当websocket连接关闭时，关闭sc管道
			server.DelConn <- c.UserId
			break
		}
		// 将json消息解析为Message结构体
		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			panic(err)
		}
		switch message.Type {
		case SEND:
			// 将消息发送给接收者
			if v, ok := server.Connections[message.Receiver]; ok {
				v.Sc <- msg
			}
		case CLOSE:
			// 关闭连接
			server.DelConn <- c.UserId
		default:
			fmt.Println("unknown message type")
		}
	}
}

// Upgrade 建立连接
func Upgrade(w http.ResponseWriter, r *http.Request) {
	var err error
	// 从请求路径中拿到userId
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	// 升级http协议到websocket协议
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	// 初始化连接
	var conn = &Connection{
		WsConn: ws,
		Sc:     make(chan []byte),
		UserId: uint64(id),
	}
	// 将连接加入到server中
	server.Connections[uint64(id)] = conn
	// 启动连接的读写协程
	go conn.writer()
	go conn.reader()
}
