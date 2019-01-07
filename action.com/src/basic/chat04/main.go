package main

import (
	"encoding/base64"
	"fmt"
)

//Base64编码
func main() {

	msg := "Hello World"
	//编码
	encodeMsg := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Base64 Encode:", encodeMsg)
	//解码
	if decodeMsg, err := base64.StdEncoding.DecodeString(encodeMsg); err != nil {
		fmt.Println("Decode Base64Text Error:", err)
	} else {
		fmt.Println("Decode Msg:", string(decodeMsg))
	}

}
