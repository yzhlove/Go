package main

import "fmt"

//空map与未初始化的map

func main() {

	//空map
	emptyMap := map[string]string{}
	emptyMap["red"] = "This Red!"
	emptyMap["green"] = "This Green!"
	emptyMap["blue"] = "This sky is blue!"

	fmt.Printf("%+v \n",emptyMap)


	//nil map 不允许进行存储
	var colorsMap map[string]string
	//colorsMap["sky"] = "sky need star!"

	fmt.Printf("%+v \n",colorsMap)

}
