package main

// for 与复合类型

func main() {

	data := [3]string{"a","b","c"}

	for i,s :=  range data {
		println("i = ",i," s = ",s)
	}

	println("-------------------")

	for i := range data {
		println("i = ",i,"data[i] = ",data[i])
	}

	println("-------------------")

	for _,s := range data {
		println(" value = ",s)
	}


}

