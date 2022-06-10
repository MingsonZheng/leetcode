package main

import "fmt"

// 剑指 Offer 06. 从尾到头打印链表
// 解法 3：自身递归

func main() {
	array := []int{1, 3, 2}
	head := initial(array)
	travel(head.Next)
	fmt.Println()
	head = head.Next
	fmt.Println(reversePrint(head))
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	subResult := reversePrint(head.Next)
	// 每次都重新初始化，不然 leetcode 里上一个测试用例的值会覆盖下一个测试用例的 result
	result := make([]int, len(subResult)+1)
	for i := 0; i < len(subResult); i++ {
		result[i] = subResult[i]
	}
	result[len(result)-1] = head.Val
	return result
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
