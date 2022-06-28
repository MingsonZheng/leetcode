package main

import "fmt"

// 114. 二叉树展开为链表

func main() {
	array := []int{1, 2, 5, 3, 4, -1, 6}
	root := MakeTree(array)
	fmt.Println(PreorderTraversal(root))
	flatten(root)
	fmt.Println(PreorderTraversal(root))
}

var dummyHead *TreeNode
var tail *TreeNode

func flatten(root *TreeNode) {
	dummyHead = &TreeNode{}
	tail = dummyHead
	preorder2(root)
}

func preorder2(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	tail.Right = root
	tail = root
	root.Left = nil
	preorder2(left)
	preorder2(right)
}
