package main

import (
	"flag"
	"fmt"
	"redis_examp/chat30/config"
	"redis_examp/chat30/opt"
)

var path = flag.String("path", "lcm", "[path]")

func main() {
	flag.Parse()
	Init()
	opt.Init()
}

func Init() {
	fmt.Println("path = ", *path)
	config.Path = *path
	fmt.Println("config path = ", config.Path)
}
