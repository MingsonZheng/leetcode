package main

import "fmt"

// 148. 排序链表
// 解法 5：快排，参考数组快排

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
	array = []int{}
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
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	linkSubQuickSort(head, tail)
	return head
}

func linkSubQuickSort(p, r *ListNode) {
	if p == r {
		return
	}
	preq, q := linkPartition(p, r)
	if q != r {
		linkSubQuickSort(q.Next, r)
	}
	if q != p {
		linkSubQuickSort(p, preq)
	}
	return
}

func linkPartition(p, r *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode = nil
	i := p
	j := p
	for j != r {
		if j.Val < r.Val {
			tmp := j.Val
			j.Val = i.Val
			i.Val = tmp
			pre = i
			i = i.Next
		}
		j = j.Next
	}
	tmp := i.Val
	i.Val = r.Val
	r.Val = tmp
	return pre, i
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
