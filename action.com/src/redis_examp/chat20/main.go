package main

import (
	"log"
	"os/exec"

	"fmt"
)

func main() {

	cmd := exec.Command("sleep", "5")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Printf("waiting...")
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok.")
}
