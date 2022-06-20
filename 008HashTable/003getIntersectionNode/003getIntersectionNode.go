package main

import "fmt"

// 160. 相交链表

func main() {
	array := []int{4, 1, 8, 4, 5}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	array2 := []int{5, 6, 1}
	head2 := initial(array2)
	travel(head2.Next)
	fmt.Println()
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// go 里面没有 hashset，用 map 代替
	hashTable := make(map[*ListNode]bool)
	p := headA
	for p != nil {
		hashTable[p] = true
		p = p.Next
	}
	p = headB
	for p != nil {
		if hashTable[p] {
			return p
		}
		p = p.Next
	}
	return nil
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
