package main

//单向通道不允许逆序操作
// close不允许接收通道

func main() {

	// c := make(chan int, 2)

	// var send chan<- int = c
	// var recv <-chan int = c

	// <-send	//发送通道不允许用于接受
	// recv <- 1 //接受通道你不允许用于发送
	// close不允许用于接收端

}
