package main

import "fmt"

// 面试题 04.06. 后继者

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2, Left: node1}
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3, Left: node2, Right: node4}
	node6 := &TreeNode{Val: 6}
	root := &TreeNode{Val: 5, Left: node3, Right: node6}
	fmt.Println(preorderTraversal(root))
	root = inorderSuccessor(root, node6)
	fmt.Println(preorderTraversal(root))
}

var coming bool
var successor *TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	coming = false
	successor = nil
	inorder(root, p)
	return successor
}

func inorder(root, p *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, p)
	if successor != nil {
		return
	}
	if coming == true {
		successor = root
		coming = false
		return
	}
	if root == p {
		coming = true
	}
	inorder(root.Right, p)
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
