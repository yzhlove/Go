package main

import (
	"example/chat32/tree"
	"fmt"
)

//便利树的节点

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}

	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}

	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	fmt.Println("Travers Tree!")

	root.Traverse()

}
