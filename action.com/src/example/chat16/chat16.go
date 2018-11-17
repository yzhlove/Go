package main

import "fmt"

//判断key是否存在

func main() {

	colors := make(map[string]string)

	//方式一

	colors["hello"] = "world"
	value , exist := colors["hello"]
	if exist {
		fmt.Printf("%v \n",value)
	}

	value , exist = colors["yzh"]
	if exist {
		fmt.Printf("%v \n",value)
	} else {
		fmt.Printf("Not Found yzh! \n")
	}


	//方式二
	result := colors["hello"]
	if result != "" {
		fmt.Printf("=> %v \n",result)
	}

	result2 := colors["yzh"]
	if result2 != "" {
		fmt.Printf("==> %v \n",result2)
	} else {
		fmt.Printf("====> Not Found !")
	}

}

