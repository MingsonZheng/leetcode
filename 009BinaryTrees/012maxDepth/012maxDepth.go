package main

import "fmt"

// 559. N 叉树的最大深度

func main() {
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node3 := &Node{Val: 3}
	children3 := []*Node{node5, node6}
	node3.Children = children3
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}
	roots := []*Node{node3, node2, node4}
	root := &Node{Val: 1}
	root.Children = roots
	fmt.Println(maxDepth(root))
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	childrenMaxDepth := 0
	for i := 0; i < len(root.Children); i++ {
		depth := maxDepth(root.Children[i])
		if depth > childrenMaxDepth {
			childrenMaxDepth = depth
		}
	}
	return childrenMaxDepth + 1
}

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
