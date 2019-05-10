package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"log"
)

func main() {

	//cmd := exec.Command("tr", "a-z", "A-Z")
	cmd := exec.Command("cat", "/Users/love/redis_filter.txt")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())

}
