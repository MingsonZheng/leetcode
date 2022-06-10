package main

import "fmt"

// 剑指 Offer 25. 合并两个排序的链表

func main() {
	array := []int{1, 2, 4}
	array2 := []int{1, 3, 4}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	head2 := initial(array2)
	travel(head2.Next)
	fmt.Println()
	res := mergeTwoLists(head.Next, head2.Next)
	travel(res)
	fmt.Println()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		subHead := mergeTwoLists(l1.Next, l2)
		l1.Next = subHead
		return l1
	} else {
		subHead := mergeTwoLists(l1, l2.Next)
		l2.Next = subHead
		return l2
	}
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 初始化
func initial(array []int) *ListNode {
	head := new(ListNode) // 虚拟头节点
	tail := head
	for _, v := range array {
		tail = insertAtTail(v, tail)
	}
	return head
}

// 在链表尾部插入元素
func insertAtTail(value int, tail *ListNode) *ListNode {
	newNode := &ListNode{value, nil}
	tail.Next = newNode
	tail = newNode
	return tail
}

// 遍历
func travel(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%v，", p.Val)
		p = p.Next
	}
}
