package main

import (
	"fmt"
	"math"
)

// 543. 二叉树的直径

func main() {
	array := []int{1, 2, 3, 4, 5}
	root := MakeTree(array)
	fmt.Println(diameterOfBinaryTree(root))
}

// 转化成求树的最大高度
var result2 int

func diameterOfBinaryTree(root *TreeNode) int {
	result2 = 0
	calMaxHeight(root)
	return result2
}

func calMaxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeftHeight := calMaxHeight(root.Left)
	maxRightHeight := calMaxHeight(root.Right)
	diameter := maxLeftHeight + maxRightHeight
	if diameter > result2 {
		result2 = diameter
	}
	return int(math.Max(float64(maxLeftHeight), float64(maxRightHeight))) + 1
}
