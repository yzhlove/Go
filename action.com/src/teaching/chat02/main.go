package main

import "fmt"

func main() {
	permuation([]rune{'a', 'b', 'c', 'd'}, 0, 4)
}

func permuation(bytes []rune, m, n int) {

	fmt.Printf("m = %d n = %d bytes = %s \n", m, n, string(bytes))

	if m < n-1 {
		permuation(bytes, m+1, n)
		for i := m + 1; i < n; i++ {
			bytes[m], bytes[i] = bytes[i], bytes[m]
			permuation(bytes, m+1, n)
			bytes[m], bytes[i] = bytes[i], bytes[m]
		}
	} else {
		fmt.Printf("%s\n", string(bytes))
	}

}
