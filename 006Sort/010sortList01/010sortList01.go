package main

import "fmt"

// 148. 排序链表
// 解法 1：递归解法

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
	midNode := findMidNode(head)
	nextNode := midNode.Next
	midNode.Next = nil
	leftHead := sortList(head)
	rightHead := sortList(nextNode)
	return mergeList(leftHead, rightHead)
}

func findMidNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeList(headA, headB *ListNode) *ListNode {
	newHead := &ListNode{}
	tail := newHead
	pa := headA
	pb := headB
	for pa != nil && pb != nil {
		if pa.Val <= pb.Val {
			tail.Next = pa
			tail = tail.Next
			pa = pa.Next
		} else {
			tail.Next = pb
			tail = tail.Next
			pb = pb.Next
		}
	}
	if pa != nil {
		tail.Next = pa
	}
	if pb != nil {
		tail.Next = pb
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
