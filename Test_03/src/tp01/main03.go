package main

import "os"

func main() {

	f,err := os.Open("./main02.go")

	buf := make([]byte,1024)
	n,err := f.Read(buf)
	f.Close()

	println("outPutText:",buf)
	println("n = ",n," err = ",err)


}
