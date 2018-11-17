package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//Go MD5

func encryTion(str, token string) (string, bool) {

	ctx := md5.New()
	_, err := ctx.Write([]byte(str + token))
	if err != nil {
		return "", false
	}
	es := ctx.Sum(nil)
	return hex.EncodeToString(es), true
}

//MD5加密
func main() {

	encryString, status := encryTion("hello world", "12345")
	if status {
		fmt.Println("encryString = ", encryString)
	}

}
