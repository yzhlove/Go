package tree

import (
	"fmt"
)

//Node 节点
type Node struct {
	Value       int
	Left, Right *Node
}

//Print 打印
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//SetValue 设置节点的值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignore.")
		return
	}
	node.Value = value
}

//CreateNode 创建节点
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
