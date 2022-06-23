package main

import "fmt"

// 538. 把二叉搜索树转换为累加树

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Right: node3}
	node0 := &TreeNode{Val: 0}
	node1 := &TreeNode{Val: 1, Left: node0, Right: node2}
	node8 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7, Right: node8}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6, Left: node5, Right: node7}
	root := &TreeNode{Val: 4, Left: node1, Right: node6}
	fmt.Println(preorderTraversal(root))
	root = convertBST(root)
	fmt.Println(preorderTraversal(root))
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	inorder(root)
	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Right)
	sum += root.Val
	root.Val = sum
	inorder(root.Left)
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	preorder(root)
	return result
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}
