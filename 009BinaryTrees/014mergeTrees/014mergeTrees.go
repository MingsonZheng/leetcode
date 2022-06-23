package main

import "fmt"

// 617. 合并二叉树

func main() {
	node5 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 3, Left: node5}
	node2 := &TreeNode{Val: 2}
	root1 := &TreeNode{Val: 1, Left: node3, Right: node2}

	node4 := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 1, Right: node4}
	node7 := &TreeNode{Val: 7}
	node33 := &TreeNode{Val: 3, Right: node7}
	root2 := &TreeNode{Val: 2, Left: node1, Right: node33}

	root := mergeTrees(root1, root2)
	fmt.Println(preorderTraversal(root))
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	newNode := &TreeNode{Val: root1.Val + root2.Val}
	leftRoot := mergeTrees(root1.Left, root2.Left)
	rightRoot := mergeTrees(root1.Right, root2.Right)
	newNode.Left = leftRoot
	newNode.Right = rightRoot
	return newNode
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
