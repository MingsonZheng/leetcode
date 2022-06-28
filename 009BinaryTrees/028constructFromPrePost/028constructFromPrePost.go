package main

import "fmt"

// 889. 根据前序和后序遍历构造二叉树

func main() {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	root := constructFromPrePost(preorder, postorder)
	fmt.Println(PreorderTraversal(root))
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return myBuildTree(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func myBuildTree(pre []int, i, j int, post []int, p, r int) *TreeNode {
	if i > j {
		return nil
	}
	root := &TreeNode{Val: pre[i]}
	if i == j {
		return root
	}
	q := p
	for post[q] != pre[i+1] {
		q++
	}
	leftTreeSize := q - p + 1
	leftNode := myBuildTree(pre, i+1, i+leftTreeSize, post, p, q)
	rightNode := myBuildTree(pre, i+leftTreeSize+1, j, post, q+1, r-1)
	root.Left = leftNode
	root.Right = rightNode
	return root
}
