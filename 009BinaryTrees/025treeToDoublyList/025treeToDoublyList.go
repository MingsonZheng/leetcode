package main

import "fmt"

// 剑指 Offer 36. 二叉搜索树与双向链表

func main() {
	array := []int{4, 2, 5, 1, 3}
	root := MakeTree(array)
	fmt.Println(PreorderTraversal(root))
	root = treeToDoublyList(root)
	for i := 0; i < 6; i++ {
		fmt.Println(root.Val)
		root = root.Right
	}
}

var dummyHead *TreeNode
var tail *TreeNode

func treeToDoublyList(root *TreeNode) *TreeNode {
	dummyHead = &TreeNode{}
	tail = dummyHead
	if root == nil {
		return nil
	}
	inorder(root)
	tail.Right = dummyHead.Right
	dummyHead.Right.Left = tail
	return dummyHead.Right
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	root.Left = tail
	tail.Right = root
	tail = root
	inorder(root.Right)
}
