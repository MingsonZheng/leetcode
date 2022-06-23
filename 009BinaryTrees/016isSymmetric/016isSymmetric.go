package main

import "fmt"

// 101. 对称二叉树

func main() {
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2, Left: node3, Right: node4}
	node44 := &TreeNode{Val: 4}
	node33 := &TreeNode{Val: 3}
	node22 := &TreeNode{Val: 2, Left: node44, Right: node33}
	root := &TreeNode{Val: 1, Left: node2, Right: node22}
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return r_isSymmetric(root.Left, root.Right)
}

func r_isSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return r_isSymmetric(p.Right, q.Left) && r_isSymmetric(p.Left, q.Right)
	}
	return false
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
