package main

import "fmt"

// 589. N 叉树的前序遍历

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
	fmt.Println(preorder(root))
}

var result []int

func preorder(root *Node) []int {
	result = make([]int, 0)
	r_preorder(root)
	return result
}

func r_preorder(root *Node) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	for i, _ := range root.Children {
		r_preorder(root.Children[i])
	}
}

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
