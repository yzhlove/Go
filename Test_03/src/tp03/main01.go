package main


func hello() {
	println("hello world")
}

func exec(f func()) {
	f()
}

func main() {

	f := hello
	exec(f)
}

