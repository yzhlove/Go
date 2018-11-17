package main

type X int

func (x *X) inc() {
	*x++
}

func main()  {


	var x X
	x.inc()
	println("x = ",x)

}
