package main

import "fmt"

func main() {
	array := []int{1, 1, 2}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	res := deleteDuplicates(head.Next)
	travel(res)
	fmt.Println()
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: -111, Next: nil} // 虚拟头节点
	tail := newHead
	p := head
	for p != nil {
		tmp := p.Next
		if p.Val != tail.Val {
			tail.Next = p
			tail = p
			p.Next = nil
		}
		p = tmp
	}
	return newHead.Next
}

//// ListNode Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//// 初始化
//func initial(array []int) *ListNode {
//	head := new(ListNode) // 虚拟头节点
//	tail := head
//	for _, v := range array {
//		tail = insertAtTail(v, tail)
//	}
//	return head
//}
//
//// 在链表尾部插入元素
//func insertAtTail(value int, tail *ListNode) *ListNode {
//	newNode := &ListNode{value, nil}
//	tail.Next = newNode
//	tail = newNode
//	return tail
//}
//
//// 遍历
//func travel(head *ListNode) {
//	p := head
//	for p != nil {
//		fmt.Printf("%v，", p.Val)
//		p = p.Next
//	}
//}
