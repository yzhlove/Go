package main

import (
	"fmt"
	"os/exec"
)

//使用command

func main() {

	cmd := exec.Command("/bin/bash", "-c", "pwd")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(output))

}
