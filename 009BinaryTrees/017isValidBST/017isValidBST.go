package main

import "fmt"

// 98. 验证二叉搜索树

func main() {
	node3 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node4 := &TreeNode{Val: 4, Left: node3, Right: node6}
	node1 := &TreeNode{Val: 1}
	root := &TreeNode{Val: 5, Left: node1, Right: node4}
	fmt.Println(isValidBST(root))
}

var isValid bool

func isValidBST(root *TreeNode) bool {
	isValid = true
	if root == nil {
		return true
	}
	dfs(root)
	return isValid
}

func dfs(root *TreeNode) []int {
	min := root.Val
	max := root.Val
	if root.Left != nil {
		leftMinMax := dfs(root.Left)
		if isValid == false {
			return nil
		}
		if leftMinMax[1] >= root.Val {
			isValid = false
			return nil
		}
		min = leftMinMax[0]
	}
	if root.Right != nil {
		rightMinMax := dfs(root.Right)
		if isValid == false {
			return nil
		}
		if rightMinMax[0] <= root.Val {
			isValid = false
			return nil
		}
		max = rightMinMax[1]
	}
	return []int{min, max}
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
