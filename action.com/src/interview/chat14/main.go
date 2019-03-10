package main

import "fmt"

type student struct {
	Name string
}

func Info(v interface{}) {
	switch _ := v.(type) {
	case *student, student:
		fmt.Println(v.(*student).Name)
	}
}

func main() {

}
