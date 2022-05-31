package main

import "fmt"

func main() {
	array := []int{1, 2, 2, 1}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	fmt.Println(isPalindrome(head.Next))
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	midNode := findMidNode(head)
	rightHalfHead := reverseList(midNode.Next)
	p := head
	q := rightHalfHead
	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func findMidNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = newHead
		newHead = p
		p = tmp
	}
	return newHead
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
