package main

import (
	"fmt"
	"math"
)

// 剑指 Offer 55 - II. 平衡二叉树

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: node9, Right: node20}
	fmt.Println(isBalanced(root))
}

var balanced bool

func isBalanced(root *TreeNode) bool {
	balanced = true
	height(root)
	return balanced
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if balanced == false {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		balanced = false
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
