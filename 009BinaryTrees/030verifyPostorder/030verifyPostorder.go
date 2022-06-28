package main

import "fmt"

// 剑指 Offer 33. 二叉搜索树的后序遍历序列

func main() {
	postorder := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder(postorder))
	postorder = []int{1, 3, 2, 6, 5}
	fmt.Println(verifyPostorder(postorder))
}

func verifyPostorder(postorder []int) bool {
	return myVerifyPostorder(postorder, 0, len(postorder)-1)
}

func myVerifyPostorder(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	// postorder[j] 是根节点，先分离出左子树 [i, k-1]
	k := i
	for k < j && postorder[k] < postorder[j] {
		k++
	}
	// 验证 [k, j-1] 满足右子树的要求，都大于 postorder[j]
	p := k
	for p < j {
		if postorder[p] < postorder[j] {
			return false
		}
		p++
	}
	// 递归验证左右子树是否满足 BST 的要求
	leftValid := myVerifyPostorder(postorder, i, k-1)
	if leftValid == false {
		return false
	}
	rightValid := myVerifyPostorder(postorder, k, j-1)
	return rightValid
}
