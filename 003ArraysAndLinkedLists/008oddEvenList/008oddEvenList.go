package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	res := oddEvenList2(head.Next)
	travel(res)
}

// 328. 奇偶链表
// 直接在原链表上操作
func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	podd := odd
	peven := even
	for podd.Next != nil && podd.Next.Next != nil {
		podd.Next = podd.Next.Next
		podd = podd.Next
		peven.Next = peven.Next.Next
		peven = peven.Next
	}
	podd.Next = even
	return odd
}

// 使用新的链表存储
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead := new(ListNode)
	oddTail := oddHead
	evenHead := new(ListNode)
	evenTail := evenHead
	p := head
	count := 1
	for p != nil {
		tmp := p.Next
		if count%2 == 1 { // 奇数
			p.Next = nil
			oddTail.Next = p
			oddTail = p
		} else { // 偶数
			p.Next = nil
			evenTail.Next = p
			evenTail = p
		}
		count++
		p = tmp
	}
	oddTail.Next = evenHead.Next
	return oddHead.Next
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
