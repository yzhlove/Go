package opt

import (
	"fmt"
	"redis_examp/chat30/config"
)

func Init() {
	fmt.Println("[opt] opt path = ", config.Path)
}
