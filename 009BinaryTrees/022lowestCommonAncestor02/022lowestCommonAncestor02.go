package main

import "fmt"

// 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
// 解法 2：递归

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
	if p == root || q == root || (p.Val < root.Val && root.Val < q.Val) || (q.Val < root.Val && root.Val < p.Val) {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
