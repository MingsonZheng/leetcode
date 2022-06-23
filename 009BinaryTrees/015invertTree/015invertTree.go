package main

import "fmt"

// 226. 翻转二叉树

func main() {
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node1, Right: node3}
	node6 := &TreeNode{Val: 6}
	node9 := &TreeNode{Val: 9}
	node7 := &TreeNode{Val: 7, Left: node6, Right: node9}
	root := &TreeNode{Val: 4, Left: node2, Right: node7}
	fmt.Println(preorderTraversal(root))
	root = invertTree(root)
	fmt.Println(preorderTraversal(root))
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	leftNode := invertTree(root.Left)
	rightNode := invertTree(root.Right)
	root.Right = leftNode
	root.Left = rightNode
	return root
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
