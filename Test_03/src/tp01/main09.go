package main

// go枚举

type color byte

const (
	black color = iota
	red
	blue
	yellow
)

func show(c color) {
	println("color = ",c)
}


func main() {

	show(black)
	show(red)
	show(blue)
	show(yellow)

}

