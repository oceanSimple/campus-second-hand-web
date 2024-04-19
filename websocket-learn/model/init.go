package model

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var server = Server{
	Connections: make(map[uint64]*Connection),
	DelConn:     make(chan uint64),
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func GetServer() *Server {
	return &server
}
