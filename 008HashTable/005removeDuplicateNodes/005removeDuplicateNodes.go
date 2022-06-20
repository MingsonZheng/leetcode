package main

import "fmt"

// 面试题 02.01. 移除重复节点

func main() {
	array := []int{1, 2, 3, 3, 2, 1}
	head := initial(array)
	head = head.Next
	travel(head)
	fmt.Println()
	removeDuplicateNodes(head)
	travel(head)
	fmt.Println()

	array = []int{1, 1, 1, 1, 2}
	head = initial(array)
	head = head.Next
	travel(head)
	fmt.Println()
	removeDuplicateNodes(head)
	travel(head)
	fmt.Println()
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	set := make(map[int]bool)
	newHead := &ListNode{}
	tail := newHead
	p := head
	for p != nil {
		tmp := p.Next
		if !set[p.Val] {
			set[p.Val] = true
			tail.Next = p
			tail = p
			tail.Next = nil
		}
		p = tmp
	}
	return newHead.Next
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
