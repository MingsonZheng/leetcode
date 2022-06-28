package main

import "fmt"

// 236. 二叉树的最近公共祖先

func main() {
	array := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root := MakeTree(array)
	p := root.Left
	q := root.Right
	fmt.Println(PreorderTraversal(root))
	root = lowestCommonAncestor(root, p, q)
	fmt.Println(PreorderTraversal(root))
}

var lca *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca = nil
	traverse(root, p, q)
	return lca
}
func traverse(root, p, q *TreeNode) int {
	if root == nil {
		return 0
	}
	leftContains := traverse(root.Left, p, q)
	if lca != nil { // 提前退出
		return 2
	}
	rightContains := traverse(root.Right, p, q)
	if lca != nil { // 提前退出
		return 2
	}
	rootContains := 0
	if root == p || root == q {
		rootContains = 1
	}
	if rootContains == 0 && leftContains == 1 && rightContains == 1 {
		lca = root
	}
	if rootContains == 1 && (leftContains == 1 || rightContains == 1) {
		lca = root
	}
	return leftContains + rightContains + rootContains
}
