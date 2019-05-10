package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("/bin/bash", "-c", "sleep 5")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Println("ok")

}
