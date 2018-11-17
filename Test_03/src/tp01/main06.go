package main


// 常量的定义以及使用

const cx, cy int = 123,0x22
const cstr = "what are you doing"
const cc = '我'

const (
	ci,cf = 1,0.234		// int ,float64(默认)
)

func main() {

	const xc = 1.23
	println("xc = ",xc)

	const yc = 1.23

	{
		const xv = "abc"
		println(xv)
	}

}

