package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"websocket-learn/model"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws/{id}", model.Upgrade)
	go model.GetServer().Run()
	//启动http服务
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
