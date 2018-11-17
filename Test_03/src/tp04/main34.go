package main

// 对字典进行迭代，每次顺序都不一样

// len返回map键值对的数量
// cap不支持map

func main() {

	m := make(map[string]int)

	for i := 0;i < 10;i++ {
		m[string('a' + i)] = i
	}

	for i:= 0;i < 4;i++ {
		for k,v := range m {
			print(k,":",v," ")
		}
		println()
	}

}
