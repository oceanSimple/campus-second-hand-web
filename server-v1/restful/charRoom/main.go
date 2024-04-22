package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server-v1/restful/charRoom/component"
	"strconv"
)

// hub 用于管理所有的连接
var hub = component.NewHub()

// upgrader 用于将http协议升级到websocket协议
var upgrader = component.GetUpgrader()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws/{id}", Upgrade)
	go hub.Run()
	//启动http服务
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}

// Upgrade 建立连接
func Upgrade(w http.ResponseWriter, r *http.Request) {
	var err error
	// 从请求路径中拿到userId
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("请求参数不是数字", err)
		return
	}
	// 升级http协议到websocket协议
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级协议失败", err)
	}
	// 初始化连接
	var client = component.NewClient(hub, uint64(id), ws)
	hub.Register <- client

	// 启动连接的读写协程
	go client.WritePump()
	go client.ReadPump()
}
