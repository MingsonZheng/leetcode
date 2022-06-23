package main

import "fmt"

// 剑指 Offer 54. 二叉搜索树的第k大节点

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2, Left: node1}
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3, Left: node2, Right: node4}
	node6 := &TreeNode{Val: 6}
	root := &TreeNode{Val: 5, Left: node3, Right: node6}
	k := 3
	fmt.Println(kthLargest(root, k))
}

var count int
var result int

func kthLargest(root *TreeNode, k int) int {
	count = 0
	result = 0
	inorder(root, k)
	return result
}

func inorder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	inorder(root.Right, k)
	if count >= k {
		return
	}
	count++
	if count == k {
		result = root.Val
		return
	}
	inorder(root.Left, k)
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
