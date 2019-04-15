package main

import (
	"fmt"
	"os/exec"
)

//Command 测试

//错误

func main() {

	cmd := exec.Command("pwd")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed" + err.Error())
		return
	}
	if bytes, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		fmt.Println("Execute Command finished.", string(bytes))
	}

}
