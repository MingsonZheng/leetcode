package main

import "fmt"

// 21. 合并两个有序链表

func main() {
	array := []int{1, 2, 4}
	head := initial(array)
	head = head.Next
	travel(head)
	fmt.Println()
	array2 := []int{1, 3, 4}
	head2 := initial(array2)
	head2 = head2.Next
	travel(head2)
	fmt.Println()
	head3 := mergeTwoLists(head, head2)
	travel(head3)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	p1 := list1
	p2 := list2
	head := &ListNode{} // 虚拟头节点
	tail := head
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			tail.Next = p1
			tail = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			tail = p2
			p2 = p2.Next
		}
	}
	// 如果 p1 还没处理完，就把剩下的直接接到 tail 后面
	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}
	return head.Next
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
