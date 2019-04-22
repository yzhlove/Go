package main

import (
	"os/exec"

	"fmt"
)

//aof

func main() {

	/*redis-server","--appendonly","yes","--save","600","1"*/
	args := []string{"run", "-d", "-p", "6379:6379", "-v", "/Users/love/DockerTest/Master:/data", "redis", "redis-server", "--appendonly", "yes"}

	if out, err := exec.Command("docker", args...).Output(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}

	fmt.Println("Done .")
}
