package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(server chan string, req string) (string, error) {
	server <- req

	select {
	case ack := <-server:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("request time out")
	}
}

func RPCServer(client chan string) {
	for {
		resp := <-client
		fmt.Println("server received:", resp)
		time.Sleep(time.Second * 2)
		client <- "successful"
	}
}

func main() {

	remote := make(chan string)
	go RPCServer(remote)

	if recv, err := RPCClient(remote, "hello"); err != nil {
		fmt.Printf("Err: %s \n", err.Error())
	} else {
		fmt.Printf("server result:%s \n", recv)
	}

}
