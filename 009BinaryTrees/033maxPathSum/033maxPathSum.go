package main

import "fmt"

// 124. 二叉树中的最大路径和

func main() {
	array := []int{-10, 9, 20, -1, -1, 15, 7}
	root := MakeTree(array)
	fmt.Println(maxPathSum(root))
}

var result2 int

func maxPathSum(root *TreeNode) int {
	result2 = -1001
	dfs(root)
	return result2
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxPath := dfs(root.Left)
	rightMaxPath := dfs(root.Right)
	max := root.Val
	if leftMaxPath > 0 {
		max += leftMaxPath
	}
	if rightMaxPath > 0 {
		max += rightMaxPath
	}
	if max > result2 {
		result2 = max
	}
	ret := root.Val
	if ret < leftMaxPath+root.Val {
		ret = leftMaxPath + root.Val
	}
	if ret < rightMaxPath+root.Val {
		ret = rightMaxPath + root.Val
	}
	return ret
}
