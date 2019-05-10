package main

import (
	"fmt"
	"os"
)

const (
	filePath = "./maze.txt"
)

//readFile 读取迷宫
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

//showMaze 显示迷宫
func showMaze(data [][]int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Printf("%-3d ", data[i][j])
		}
		fmt.Println()
	}
}

//point 按 i,j 取更符合编程时的坐标系
type point struct {
	i, j int
}

func (p point) add(tp point) point {
	return point{p.i + tp.i, p.j + tp.j}
}

func (p point) at(graph [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(graph) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(graph[p.i]) {
		return 0, false
	}
	return graph[p.i][p.j], true
}

//上 右 下 左
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(graph [][]int, start, end point) [][]int {

	row := len(graph)
	col := len(graph[0])
	setps := make([][]int, row)
	for i := 0; i < row; i++ {
		setps[i] = make([]int, col)
	}
	//初始化队列
	queue := []point{start}
	var distance int
	var ok bool
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			if distance, ok = next.at(graph); !ok || distance == 1 {
				continue
			}
			if distance, ok = next.at(setps); !ok || distance != 0 {
				continue
			}
			if next == start {
				continue
			}
			distance, _ = cur.at(setps)
			setps[next.i][next.j] = distance + 1
			queue = append(queue, next)
		}
	}
	return setps
}

func main() {

	tmp := readFile(filePath)
	showMaze(tmp)

	fmt.Println()

	seps := walk(tmp, point{0, 0}, point{len(tmp) - 1, len(tmp[0]) - 1})
	showMaze(seps)
}
