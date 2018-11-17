package main

import (
	"flag"
	"fmt"
)

//flag使用

func main() {

	var host string
	var port int

	isParse := flag.Parsed()
	fmt.Println("isFlag:", isParse)

	//flag 参数 [ 变量指针，参数名，默认值，帮助信息]
	flag.StringVar(&host, "host", "www.baidu.com", "please input URL")
	flag.IntVar(&port, "port", 80, "please input port")
	flag.Parse()
	isParse = flag.Parsed()
	fmt.Println("isFlag:", isParse)

	fmt.Println("---------- value ----------")
	fmt.Println("host:", host, " port:", port)

	fmt.Println("---------- args start ----------")
	// visit只包含已经设置值的
	flag.Visit(func(fn *flag.Flag) {
		fmt.Println(fn.Name, fn.Value, fn.Usage, fn.DefValue)
	})
	fmt.Println("---------- args end ----------")

	fmt.Println("---------- args start ----------")
	flag.VisitAll(func(fn *flag.Flag) {
		fmt.Println(fn.Name, fn.Value, fn.Usage, fn.DefValue)
	})
	fmt.Println("---------- args end ----------")

	flag.PrintDefaults()

	//非flag参数的个数
	fmt.Printf("Nargs:", flag.NArg)
	//flag参数的个数
	fmt.Printf("Fargs:", flag.NFlag)

}
