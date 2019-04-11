package main

type User struct {
	Name string
}

func GetName(index int) string {
	lists := []string{"lcm", "xjj", "xyj", "hxy", "fyb"}
	return lists[index]
}

func main() {

	//value := rand.Intn(5)
	//
	//switch GetName(value) {
	////u := User{Name:name}
	//case "lcm":
	//	fmt.Println("lcm", u)
	//case "xjj":
	//	fmt.Println("lcm", u)
	//case "xyj":
	//	fmt.Println("lcm", u)
	//case "hxy":
	//	fmt.Println("lcm", u)
	//case "fyb":
	//	fmt.Println("lcm", u)
	//}

}
