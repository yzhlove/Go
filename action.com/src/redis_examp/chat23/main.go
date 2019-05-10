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

	out, err := exec.Command("docker", "ps", "-a").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

}

func Init() error {

	return nil
}

func cmdRun(name, flag string, args ...string) error {

	return nil
}
