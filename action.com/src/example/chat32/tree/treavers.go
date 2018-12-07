package tree

import (
	"fmt"
)

//Traverse 便利
func (node *Node) Traverse() {
	node.TraversFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

//TraversFunc 便利函数
func (node *Node) TraversFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraversFunc(f)
	f(node)
	node.Right.TraversFunc(f)
}

//TraversWithChannel 利用通道便利
func (node *Node) TraversWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraversFunc(func(n *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
