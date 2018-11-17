package main

import "strconv"

//  闭包个性

func testFunc() []func() {

	var funcs []func()

	for i := 0;i < 5;i++ {
		var str = strconv.FormatInt(int64(i),10) + ":"
		funcs = append(funcs, func() {
			println(str,&i,i)
		})
	}

	return funcs

}



func main() {

	for _,f := range testFunc() {
		f()
	}


}
