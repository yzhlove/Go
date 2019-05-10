package main

import (
	"fmt"
	"hash/crc32"
)

func main() {

	values := []string{"lcm", "xyj", "yzh", "fyb", "xjj", "hxy", "cwq", "zwl", "dyy", "wrs"}
	for _, v := range values {

		index := crc32.ChecksumIEEE([]byte(v))

		fmt.Printf("%v %v \n", index, index%10)

	}

}
