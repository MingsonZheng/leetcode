package main

import "fmt"

// 590. N 叉树的后序遍历

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
	fmt.Println(postorder(root))
}

var result []int

func postorder(root *Node) []int {
	result = make([]int, 0)
	r_postorder(root)
	return result
}

func r_postorder(root *Node) {
	if root == nil {
		return
	}
	for i, _ := range root.Children {
		r_postorder(root.Children[i])
	}
	result = append(result, root.Val)
}

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
