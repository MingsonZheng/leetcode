package main

import "fmt"

// 94. 二叉树的中序遍历

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	root := &TreeNode{Val: 1, Right: node2}
	fmt.Println(inorderTraversal(root))
}

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	inorder(root)
	return result
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	result = append(result, root.Val)
	inorder(root.Right)
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
