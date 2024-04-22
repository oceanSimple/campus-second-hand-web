package component

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func GetUpgrader() websocket.Upgrader {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	return upgrader
}
