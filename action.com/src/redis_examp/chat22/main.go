package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("/bin/bash", "-c", "docker info").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
