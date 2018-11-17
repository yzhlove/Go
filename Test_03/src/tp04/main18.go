package main

// 指针与数组

func main() {

	a := [...]int{1,2}
	p := &a

	p[1] += 100

	println("p[1] = ",p[1])

}


