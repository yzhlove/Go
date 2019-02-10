package main

import (
	"fmt"
	"os"
)

const (
	filePath = "./maze.txt"
)

func readFile(path string) [][]int {
	var (
		file          *os.File
		err           error
		row, col, num int
	)
	if file, err = os.Open(path); err != nil {
		panic(err)
	}
	if num, err = fmt.Fscanf(file, "%d %d", &row, &col); err != nil || num != 2 {
		panic(err)
	}
	//初始化二维迷宫
	data := make([][]int, row)
	for i := 0; i < row; i++ {
		data[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			_, _ = fmt.Fscanf(file, "%d", &data[i][j])
		}
	}
	return data
}

func main() {

	tmp := readFile(filePath)
	fmt.Println(tmp)

}
