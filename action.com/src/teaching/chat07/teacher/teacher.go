package teacher

import (
	"fmt"
	"teaching/chat07/base"
)

type Teacher struct{}

func (t *Teacher) Info() {
	fmt.Println(" <--- teacher --->")
}

func init() {
	base.Register("teacher", func() base.User {
		return new(Teacher)
	})
}
