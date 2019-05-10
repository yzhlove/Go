package main

import (
	"fmt"
	"os/exec"
)

/*
docker images | grep redis
docker pull redis
docker rm saved
docker run -d -p 6380:6379 --name saved redis
docker run -d -p 6379:6379 --link saved:saved -v /Develop/Docker/Volumes/Redis:/data redis
*/

func main() {
	_ = runMaster()
}

func runMaster() error {
	//cmd := `docker run -d -p 6379\:6379 --name saved redis`
	if out, err := exec.Command("docker", "run", "-d", "-p", "6379:6379", "--name", "saved", "redis").Output(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
	return nil
}
