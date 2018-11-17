package main

func main() {

	m := map[string]int {
		"abc":123,
	}

	key := []byte("abc")
	x , ok := m[string(key)]

	println(x,ok)

}
