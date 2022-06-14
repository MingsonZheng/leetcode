package main

import (
	"fmt"
	"math"
)

// 147. 对链表进行插入排序

func main() {
	array := []int{4, 2, 1, 3}
	head := initial(array)
	head = head.Next
	travel(head)
	fmt.Println()
	travel(insertionSortList(head))
	fmt.Println()
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 存储已经排序好的节点
	newHead := &ListNode{Val: math.MaxInt32, Next: nil}
	// 遍历节点
	p := head
	for p != nil {
		tmp := p.Next
		// 寻找 p 节点插入的位置，插入到哪个节点后面
		q := newHead                               // 初始化值
		for q.Next != nil && q.Next.Val <= p.Val { // 循环条件结束
			q = q.Next
		}
		// 将 p 节点插入
		p.Next = q.Next
		q.Next = p
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
