package main

// switch

func main() {

	a,b,c,x := 1,2,3,2

	switch x {
	case a,b:
		println("a | b")
	case c:
		println("c")
	case 4:
		println("d")
	default:
		println("z")
	}

	println("------------------------")

	switch value := 5;value {
	default:				// go编译器回确保最后执行default
		value += 100
		println("value = " , value)
	case 5:
		value += 50
		println("value = ",value)

	}

	println("------------------------")

	// go switch 会在梅哥case 后面自动break,如需继续执行，可使用 fallthroungh

	switch expr := 5;expr {
	default:
		println("expr = " ,expr )
	case 5:
		expr += 10
		println("expr = " ,expr )
		fallthrough
	case 6:
		expr += 20
		println("expr = " ,expr )
	}

}