package main

import "os"

//文件测试

func main() {

	var (
		err  error
		file *os.File
		mode = os.O_RDWR | os.O_CREATE | os.O_APPEND
	)
	if file, err = os.OpenFile("./test.txt", mode, 0666); err != nil {
		panic(err)
	}

	defer file.Close()
	_, _ = file.Write([]byte("what are you doing!\r\n"))

}
