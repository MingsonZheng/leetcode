package main

import "fmt"

func main() {
	array := []int{2, 4, 3}
	array2 := []int{5, 6, 4}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	head2 := initial(array2)
	travel(head2.Next)
	fmt.Println()
	res := addTwoNumbers(head.Next, head2.Next)
	travel(res)
	fmt.Println()
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	head := new(ListNode) // 虚拟头节点
	tail := head
	carry := 0
	for p1 != nil || p2 != nil {
		sum := 0
		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}
		if carry != 0 {
			sum += carry
		}
		tail.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		tail = tail.Next
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head.Next
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
