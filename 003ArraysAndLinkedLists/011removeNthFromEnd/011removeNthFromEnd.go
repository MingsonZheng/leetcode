package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	res := removeNthFromEnd(head.Next, 2)
	travel(res)
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 遍历 1：fast 先到第 n 个节点
	fast := head
	count := 0
	for fast != nil {
		count++
		if count == n {
			break
		}
		fast = fast.Next
	}
	if fast == nil { // 链表节点不够 k
		return head
	}
	// 遍历 2
	slow := head
	var pre *ListNode = nil
	for fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	// 删除倒数第 n 个节点
	if pre == nil { // 头节点是倒数第 n 个节点
		head = head.Next
	} else {
		pre.Next = slow.Next
	}
	return head
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
