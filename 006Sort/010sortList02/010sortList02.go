package main

import "fmt"

// 148. 排序链表
// 解法 2：非递归解法（不好理解）

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
	n := len(head)
	step := 1
	for step < n {
		newHead := &ListNode{}
		tail := newHead
		p := head
		for p != nil {
			// [p, q]
			q := p
			count := 1
			for q != nil && count < step {
				q = q.Next
				count++
			}
			if q == nil || q.Next == nil { // 这一轮合并结束了
				tail.Next = p
				break
			}
			// [q+1, r]
			r := q.Next
			count = 1
			for r != nil && count < step {
				r = r.Next
				count++
			}
			// 保存下一个 step 的起点
			var tmp *ListNode = nil
			if r != nil {
				tmp = r.Next
			}
			// merge [p, q] [q+1, r]
			tail.Next, tail = merge(p, q, r)
			p = tmp
		}
		head = newHead.Next
		step *= 2
	}
	return head
}

func len(head *ListNode) int {
	if head == nil {
		return 0
	}
	n := 1
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	return n
}

func merge(p, q, r *ListNode) (*ListNode, *ListNode) {
	newHead := &ListNode{}
	tail := newHead
	pa := p
	pb := q.Next
	q.Next = nil
	if r != nil {
		r.Next = nil
	}
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
		tail = q
	}
	if pb != nil {
		tail.Next = pb
		tail = r
	}
	return newHead.Next, tail
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
