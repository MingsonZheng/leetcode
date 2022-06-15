package main

import "fmt"

// 148. 排序链表
// 解法 3：链表冒泡排序（交换节点，有点费劲）

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
	p := head
	n := 0
	for p != nil {
		p = p.Next
		n++
	}
	for i := 0; i < n; i++ {
		q := head
		var pre *ListNode = nil
		for j := 0; j < n-i-1; j++ {
			// 交换 q 和 q.next；先删除 q.next，并记录为 tmp；再把 tmp 插入到 pre 和 q 之间
			if q.Val > q.Next.Val {
				// 交换结束后，q 已经位于 q.next；因此不必另加一句 q = q.Next
				if pre == nil { // q 是头节点
					pre = q.Next
					tmp := q.Next
					q.Next = q.Next.Next
					tmp.Next = q
					head = tmp
				} else {
					tmpPre := pre
					// q 和 q.next 交换，交换完后，q.next 会成为 q 的 pre，因此这里提前记录下
					pre = q.Next
					tmp := q.Next
					q.Next = q.Next.Next
					tmp.Next = q
					tmpPre.Next = tmp
				}
			} else {
				pre = q
				q = q.Next
			}
		}
	}
	return head
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
