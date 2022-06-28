package main

var result []int

func PreorderTraversal(root *TreeNode) []int {
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
