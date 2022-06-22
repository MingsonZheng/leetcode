package main

import (
	"container/list"
	"fmt"
)

// 剑指 Offer 32 - I. 从上到下打印二叉树
// 解法 2：list 实现 queue

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: node9, Right: node20}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	queue := list.List{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		front := queue.Front()
		node := front.Value.(*TreeNode)
		result = append(result, node.Val)
		queue.Remove(front)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
