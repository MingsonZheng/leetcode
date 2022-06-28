package main

import "fmt"

// 面试题 17.12. BiNode

func main() {
	array := []int{4, 2, 5, 1, 3, -1, 6, 0}
	root := MakeTree(array)
	fmt.Println(PreorderTraversal(root))
	root = convertBiNode(root)
	fmt.Println(PreorderTraversal(root))
}

var dummyHead *TreeNode
var tail *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	dummyHead = &TreeNode{}
	tail = dummyHead
	inorder(root)
	return dummyHead.Right
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	tail.Right = root
	tail = root
	root.Left = nil
	inorder(root.Right)
}
