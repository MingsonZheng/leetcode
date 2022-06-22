package main

import "fmt"

// 145. 二叉树的后序遍历

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	root := &TreeNode{Val: 1, Right: node2}
	fmt.Println(postorderTraversal(root))
}

var result []int

func postorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	postorder(root)
	return result
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	result = append(result, root.Val)
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
