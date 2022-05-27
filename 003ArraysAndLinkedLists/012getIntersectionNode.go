package main

import "fmt"

func main() {
	array := []int{4, 1, 8, 4, 5}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	array2 := []int{5, 6, 1}
	head2 := initial(array2)
	//travel(head2.Next)
	//fmt.Println()
	p := head2
	for p.Next != nil {
		p = p.Next
	}
	tail2 := p
	tail2.Next = head.Next.Next.Next // array2 的尾指针指向 array 的第三个节点
	travel(head2.Next)
	fmt.Println()
	res := getIntersectionNode(head.Next, head2.Next)
	travel(res)
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 求链表 A 的长度 na
	na := 0
	pA := headA
	for pA != nil {
		na++
		pA = pA.Next
	}
	// 求链表 B 的长度 nb
	nb := 0
	pB := headB
	for pB != nil {
		nb++
		pB = pB.Next
	}

	// 先让指向长链表的指针多走 na-nb 或 nb-na 步
	pA = headA
	pB = headB
	if na >= nb {
		for i := 0; i < na-nb; i++ {
			pA = pA.Next
		}
	} else {
		for i := 0; i < nb-na; i++ {
			pB = pB.Next
		}
	}
	// 让 pA 和 pB 同步前进
	for pA != nil && pB != nil && pA != pB {
		pA = pA.Next
		pB = pB.Next
	}
	if pA == nil && pB == nil {
		return nil
	} else {
		return pA
	}
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
