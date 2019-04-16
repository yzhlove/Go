package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("/bin/bash", "-c", "ls -l")
	if out, err := cmd.Output(); err != nil {
		panic(err)
		return
	} else {
		fmt.Println("value = ", string(out))
	}

}
