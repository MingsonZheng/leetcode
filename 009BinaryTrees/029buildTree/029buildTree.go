package main

import "fmt"

// 106. 从中序与后序遍历序列构造二叉树

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println(PreorderTraversal(root))
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return myBuildTree(postorder, 0, len(postorder)-1, inorder, 0, len(inorder)-1)
}

func myBuildTree(postorder []int, i, j int, inorder []int, p, r int) *TreeNode {
	if i > j {
		return nil
	}
	root := &TreeNode{Val: postorder[j]}
	q := p
	for inorder[q] != postorder[j] {
		q++
	}
	leftTreeSize := q - p
	leftNode := myBuildTree(postorder, i, i+leftTreeSize-1, inorder, p, q-1)
	rightNode := myBuildTree(postorder, i+leftTreeSize, j-1, inorder, q+1, r)
	root.Left = leftNode
	root.Right = rightNode
	return root
}
