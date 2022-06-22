package main

import "fmt"

// 144. 二叉树的前序遍历
// 解法 2：非递归

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	root := &TreeNode{Val: 1, Right: node2}
	fmt.Println(preorderTraversal(root))
}

type SFrame struct {
	status int
	node   *TreeNode
}

var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]SFrame, 0)
	stack = append(stack, SFrame{status: 1, node: root})
	for len(stack) > 0 {
		if stack[len(stack)-1].status == 1 {
			result = append(result, stack[len(stack)-1].node.Val)
			stack[len(stack)-1].status = 2
			if stack[len(stack)-1].node.Left != nil {
				stack = append(stack, SFrame{status: 1, node: stack[len(stack)-1].node.Left})
			}
			continue
		}
		if stack[len(stack)-1].status == 2 {
			stack[len(stack)-1].status = 3
			if stack[len(stack)-1].node.Right != nil {
				stack = append(stack, SFrame{status: 1, node: stack[len(stack)-1].node.Right})
			}
			continue
		}
		if stack[len(stack)-1].status == 3 {
			stack = stack[:len(stack)-1]
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
