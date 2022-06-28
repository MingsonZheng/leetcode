package main

import "fmt"

// 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先

func main() {
	array := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	root := MakeTree(array)
	p := root.Left
	q := root.Right
	fmt.Println(PreorderTraversal(root))
	root = lowestCommonAncestor(root, p, q)
	fmt.Println(PreorderTraversal(root))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x := root
	for {
		if p.Val < x.Val && q.Val < x.Val {
			x = x.Left
		} else if p.Val > x.Val && q.Val > x.Val {
			x = x.Right
		} else {
			return x
		}
	}
}
