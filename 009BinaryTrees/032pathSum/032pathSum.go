package main

import "fmt"

// 剑指 Offer 34. 二叉树中和为某一值的路径

func main() {
	array := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}
	root := MakeTree(array)
	fmt.Println(pathSum(root, 22))
}

var result2 [][]int

func pathSum(root *TreeNode, target int) [][]int {
	result2 = make([][]int, 0)
	if root == nil {
		return result2
	}
	dfs(root, target, []int{}, 0)
	return result2
}

func dfs(root *TreeNode, sum int, path []int, pathSum int) {
	path = append(path, root.Val)
	pathSum += root.Val
	if root.Left == nil && root.Right == nil {
		if pathSum == sum {
			pathSnapshot := make([]int, len(path))
			copy(pathSnapshot, path)
			result2 = append(result2, pathSnapshot)
		}
		path = path[:len(path)-1]
		return
	}
	if root.Left != nil {
		dfs(root.Left, sum, path, pathSum)
	}
	if root.Right != nil {
		dfs(root.Right, sum, path, pathSum)
	}
	path = path[:len(path)-1]
}
