package main

import "fmt"

// 105. 从前序与中序遍历序列构造二叉树

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(PreorderTraversal(root))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return myBuildTree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func myBuildTree(preorder []int, i, j int, inorder []int, p, r int) *TreeNode {
	if i > j {
		return nil
	}
	root := &TreeNode{Val: preorder[i]}
	q := p
	for inorder[q] != preorder[i] {
		q++
	}
	leftTreeSize := q - p
	leftNode := myBuildTree(preorder, i+1, i+leftTreeSize, inorder, p, q-1)
	rightNode := myBuildTree(preorder, i+leftTreeSize+1, j, inorder, q+1, r)
	root.Left = leftNode
	root.Right = rightNode
	return root
}
