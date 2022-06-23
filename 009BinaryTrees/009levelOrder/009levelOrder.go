package main

import "fmt"

// 429. N 叉树的层序遍历

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
	fmt.Println(levelOrder(root))
}

var result [][]int

func levelOrder(root *Node) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			for j := 0; j < len(node.Children); j++ {
				queue = append(queue, node.Children[j])
			}
		}
		result = append(result, level)
	}
	return result
}

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
