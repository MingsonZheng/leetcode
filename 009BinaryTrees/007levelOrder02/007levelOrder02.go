package main

import "fmt"

// 102. 二叉树的层序遍历
// 解法 2：dfs

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: node9, Right: node20}
	fmt.Println(levelOrder(root))

	root = &TreeNode{Val: 1}
	fmt.Println(levelOrder(root))
}

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = nil
	result := make([][]int, 0)
	result = dfs(root, 0)
	return result
}

func dfs(root *TreeNode, level int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if level > len(result)-1 {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
