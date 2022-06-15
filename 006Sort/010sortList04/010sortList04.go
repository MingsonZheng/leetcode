package main

import "fmt"

// 148. 排序链表
// 解法 4：链表冒泡排序（交换值，建议就用交换值吧）

func main() {
	array := []int{4, 2, 1, 3}
	head := initial(array)
	head = head.Next
	travel(head)
	fmt.Println()
	travel(sortList(head))
	fmt.Println()
	array = []int{-1, 5, 3, 4, 0}
	head = initial(array)
	head = head.Next
	travel(head)
	fmt.Println()
	travel(sortList(head))
	fmt.Println()
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	p := head
	n := 0
	for p != nil {
		p = p.Next
		n++
	}
	for i := 0; i < n; i++ {
		q := head
		for j := 0; j < n-i-1; j++ {
			if q.Val > q.Next.Val {
				tmp := q.Val
				q.Val = q.Next.Val
				q.Next.Val = tmp
			}
			q = q.Next
		}
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
