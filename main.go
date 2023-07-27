package main

import (
	"awesomeProject/router"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//优化路由功能
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
