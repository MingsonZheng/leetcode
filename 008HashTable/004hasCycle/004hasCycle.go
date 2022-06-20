package main

import "fmt"

// 141. 环形链表

func main() {
	array := []int{3, 2, 0, -4}
	head := initial(array)
	p := head
	for p.Next != nil {
		p = p.Next
	}
	tail := p
	tail.Next = head.Next.Next // 链表尾部指向第二个元素
	fmt.Println()
	fmt.Println(hasCycle(head.Next))
}

func hasCycle(head *ListNode) bool {
	// go 里面没有 hashset，用 map 代替
	hashTable := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if hashTable[p] {
			return true
		} else {
			hashTable[p] = true
		}
		p = p.Next
	}
	return false
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
