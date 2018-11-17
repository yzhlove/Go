package opt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


type defaultGraphics struct {}

//Graphicser 视图接口
var Graphicser interface{
	ShowString() string
}

func Run() {

	inputReader := bufio.NewReader(os.Stdin)

	//创建默认视图
	dfg := new(defaultGraphics)

	for {
		fmt.Println(dfg.ShowString())
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
			break
		}
		switch input = strings.Replace(input,"\n","",-1) ; input {
		case "quit":
			fmt.Println("Done!")
			return
		case "show":
			userIfs, err :=  InitUserData()
			if err != nil {
				fmt.Println("Done!")
				return
			}
			//显示信息
			for _,userInfo := range userIfs {
				go userInfo.optUser()
			}
		case "query":
			fmt.Println("This feature is not available yet !")
		default:
			fmt.Println("Option Invalid:",input)
		}
	}
	fmt.Println("Done!")
}


func (*defaultGraphics) ShowString() string {

	var str string
	str += "+---------------------------------+\n"
	str += "| show)  显示 \n"
	str += "| query) 查询 \n"
	str += "| quit)  退出 \n"
	str += "+---------------------------------+\n"

	return str

}

