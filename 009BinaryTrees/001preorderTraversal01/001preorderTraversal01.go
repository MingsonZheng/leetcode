package main

import "fmt"

// 144. 二叉树的前序遍历
// 解法 1：递归

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	root := &TreeNode{Val: 1, Right: node2}
	fmt.Println(preorderTraversal(root))
}

var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	preorder(root)
	return result
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
