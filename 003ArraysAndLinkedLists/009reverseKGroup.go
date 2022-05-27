package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	res := reverseKGroup(head.Next, 2)
	travel(res)
}

// 25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := new(ListNode)
	tail := dummyHead
	p := head
	for p != nil {
		count := 0
		q := p
		for q != nil {
			count++
			if count == k {
				break
			}
			q = q.Next
		}
		if q == nil {
			tail.Next = p
			return dummyHead.Next
		} else {
			tmp := q.Next
			nodes := reverse(p, q)
			tail.Next = nodes[0]
			tail = nodes[1]
			p = tmp
		}
	}
	return dummyHead.Next
}

func reverse(head *ListNode, tail *ListNode) []*ListNode {
	var newHead *ListNode = nil
	p := head
	for p != tail {
		tmp := p.Next
		p.Next = newHead
		newHead = p
		p = tmp
	}
	tail.Next = newHead
	//newHead = tail
	return []*ListNode{tail, head}
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
