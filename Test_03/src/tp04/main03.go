package main

// string 不做转义处理

func main() {

	s := `line\r\n,
		line2`
	println(s,":",len(s))

	s = `what are you doing.
		Hello World,
		<html>
			When are you doing
		</html>
		`
	println(s)
}
