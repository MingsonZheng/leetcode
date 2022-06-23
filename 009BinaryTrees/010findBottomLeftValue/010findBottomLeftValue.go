package main

import "fmt"

// 513. 找树左下角的值

/*
*     1
*    / \
*   2  3
*  /  /  \
* 4  5   6
*   /
*  7
 */

func main() {
	node7 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 5, Left: node7}
	node6 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 3, Left: node5, Right: node6}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2, Left: node4}
	root := &TreeNode{Val: 1, Left: node2, Right: node3}
	fmt.Println(findBottomLeftValue(root))
}

func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := -1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = node.Val
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
