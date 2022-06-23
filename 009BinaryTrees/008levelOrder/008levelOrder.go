package main

import "fmt"

// 剑指 Offer 32 - III. 从上到下打印二叉树 III

/*
*    3
*   / \
*  9  20
*    /  \
*   15   7
 */

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: node9, Right: node20}
	fmt.Println(levelOrder(root))
}

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	stacks := make([][]*TreeNode, 2)
	for i := 0; i < 2; i++ {
		stacks[i] = make([]*TreeNode, 0)
	}
	turn := 0
	stacks[turn] = append(stacks[turn], root) // [[3][]]
	for len(stacks[turn]) > 0 {               // 1; 2
		curLevelNodes := make([]int, 0)
		for len(stacks[turn]) > 0 {
			treeNode := stacks[turn][len(stacks[turn])-1]       // 3; 20
			stacks[turn] = stacks[turn][:len(stacks[turn])-1]   // [[][]]; [[][9]]
			curLevelNodes = append(curLevelNodes, treeNode.Val) // [3]; [20]
			if turn == 0 {
				if treeNode.Left != nil {
					stacks[1] = append(stacks[1], treeNode.Left) // [[][9]]
				}
				if treeNode.Right != nil {
					stacks[1] = append(stacks[1], treeNode.Right) // [[][9, 20]]
				}
			} else {
				if treeNode.Right != nil {
					stacks[0] = append(stacks[0], treeNode.Right) // [[7][9]]
				}
				if treeNode.Left != nil {
					stacks[0] = append(stacks[0], treeNode.Left) // [[15][9]]
				}
			}
		}
		result = append(result, curLevelNodes) // [[3][]]
		turn = (turn + 1) % 2                  // 1
	}
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
