package main

import (
	"fmt"
	"math"
)

// 104. 二叉树的最大深度

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: node9, Right: node20}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
