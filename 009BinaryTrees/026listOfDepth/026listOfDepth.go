package main

import "fmt"

// 面试题 04.03. 特定深度节点链表

func main() {
	array := []int{1, 2, 3, 4, 5, -1, 7, 8}
	root := MakeTree(array)
	fmt.Println(PreorderTraversal(root))
	res := listOfDepth(root)
	fmt.Printf("[")
	for i := 0; i < len(res); i++ {
		listNode := res[i]
		fmt.Printf("[%d", listNode.Val)
		for listNode.Next != nil {
			fmt.Printf(", %d", listNode.Next.Val)
			listNode = listNode.Next
		}
		if i < len(res)-1 {
			fmt.Printf("], ")
		} else {
			fmt.Printf("]")
		}
	}
	fmt.Printf("]")
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return []*ListNode{}
	}
	result := make([]*ListNode, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)
	for len(queue) > 0 {
		dummyHead := &ListNode{}
		tail := dummyHead
		curLevelNum := len(queue)
		for i := 0; i < curLevelNum; i++ {
			treeNode := queue[0]
			queue = queue[1:]
			tail.Next = &ListNode{Val: treeNode.Val}
			tail = tail.Next
			if treeNode.Left != nil {
				queue = append(queue, treeNode.Left)
			}
			if treeNode.Right != nil {
				queue = append(queue, treeNode.Right)
			}
		}
		result = append(result, dummyHead.Next)
	}
	return result
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
