package main

import "fmt"

func main() {
	array := []int{3, 2, 0, -4}
	head := initial(array)
	p := head
	for p.Next != nil {
		p = p.Next
	}
	tail := p
	tail.Next = head.Next.Next // 链表尾部指向第二个元素
	fmt.Println()
	fmt.Println(hasCycle(head.Next))
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow == fast {
		return true
	}
	return false
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
