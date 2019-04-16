package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

//pipe

func main() {

	cmd := exec.Command("echo", "-n", `{"Name":"Bod","Age":32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("person: %v \n", person)
}
