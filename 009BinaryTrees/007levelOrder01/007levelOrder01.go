package main

import "fmt"

// 102. 二叉树的层序遍历

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: node9, Right: node20}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLevelNodes := make([]int, 0)
		curLevelNum := len(queue)
		for i := 0; i < curLevelNum; i++ {
			treeNode := queue[0]
			queue = queue[1:]
			curLevelNodes = append(curLevelNodes, treeNode.Val)
			if treeNode.Left != nil {
				queue = append(queue, treeNode.Left)
			}
			if treeNode.Right != nil {
				queue = append(queue, treeNode.Right)
			}
		}
		result = append(result, curLevelNodes)
	}
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
