package main

import (
	"log"
	"logging/chat11/handlers"
	"net/http"
)

//测试服务端点

func main() {
	handlers.Routes()
	log.Println("listener:Start On Port:4000")
	http.ListenAndServe(":4000", nil)
}
