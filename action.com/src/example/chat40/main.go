package main

import (
	"example/chat40/mock"
	"example/chat40/retiver"
	"fmt"
	"time"
)

//接口组合
const url = "http://www.imooc.com"

//Retriever 接口
type Retriever interface {
	Get(url string) string
}

//Poster 接口
type Poster interface {
	Post(url string, from map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

//RetrieverPoster 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another facked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("Type:%T Value:%v \n", r, r)
	fmt.Print("Type switch:")

	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *retiver.Retriver:
		fmt.Println("UaseAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever

	mockRetriver := mock.Retriver{
		Contents: "this is a fack imooc.com",
	}
	r = &mockRetriver
	inspect(r)

	r = &retiver.Retriver{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//Type assertion 类型断言
	if mockRetriver, ok := r.(*mock.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("r is not a mock retriver")
	}
	fmt.Println("Try a session with mockRetriver")
	fmt.Println(session(&mockRetriver))

}
