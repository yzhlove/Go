package main

import "fmt"

// 复合结构体

func main() {

	type user struct {
		name string
		age byte
	}

	d := [...]user {
		user{"Tome",20},		// 可以省略具体类型，如下
		{"Mary",18},
	}

	fmt.Printf("%+v\n",d)

}


/*
fmt.Printf("%#v\n",d) :
[2]main.user{main.user{name:"Tome", age:0x14}, main.user{name:"Mary", age:0x12}}

fmt.Printf("%+v\n",d) :
[{name:Tome age:20} {name:Mary age:18}]

*/