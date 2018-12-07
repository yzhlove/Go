package main

import (
	"log"
	"math/rand"
)

//fmt打印行号和

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	log.Println(rand.Intn(10))
}
