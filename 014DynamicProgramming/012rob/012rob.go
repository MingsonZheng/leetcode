package main

import (
	"fmt"
	"math"
)

// 337. 打家劫舍 III

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Right: node3}
	node1 := &TreeNode{Val: 1}
	node33 := &TreeNode{Val: 3, Right: node1}
	root := &TreeNode{Val: 3, Left: node2, Right: node33}
	fmt.Println(rob(root))
}

func rob(root *TreeNode) int {
	money := postorder(root)
	return int(math.Max(float64(money[0]), float64(money[1])))
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	leftMoney := postorder(root.Left)
	rightMoney := postorder(root.Right)
	money := make([]int, 2)
	money[0] = int(math.Max(float64(leftMoney[0]), float64(leftMoney[1]))) +
		int(math.Max(float64(rightMoney[0]), float64(rightMoney[1])))
	money[1] = (leftMoney[0] + rightMoney[0]) + root.Val
	return money
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
