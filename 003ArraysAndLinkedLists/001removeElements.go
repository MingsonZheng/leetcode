package main

import "fmt"

func main() {
	head := new(ListNode) // 虚拟头节点
	tail := head
	tail = insertAtTail(1, tail)
	tail = insertAtTail(2, tail)
	tail = insertAtTail(6, tail)
	tail = insertAtTail(3, tail)
	tail = insertAtTail(4, tail)
	tail = insertAtTail(5, tail)
	tail = insertAtTail(6, tail)
	travel(head.Next)
	fmt.Println()
	removeElements3(head, 6)
	travel(head.Next)
	fmt.Println()
}

// 203. 移除链表元素
// 遍历的结束条件：prev.Next == null
// 指针的初始值：prev := head
// 遍历的核心逻辑：
// if prev.Next.Val == val { prev.Next = prev.Next.Next }
// else { prev = prev.Next }
// 特殊情况处理：头节点、尾节点、空链表
// 引入虚拟节点：可以通过添加虚拟节点简化编程

// 解法 3：改变链表的万能写法（返回一个新的链表）
func removeElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	tail := newHead
	p := head
	for p != nil {
		tmp := p.Next
		if p.Val != val {
			p.Next = nil
			tail.Next = p
			tail = p
		}
		p = tmp
	}
	return newHead.Next
}

// 解法 2：添加虚拟头节点
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = head
	prev := newHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return newHead.Next
}

// 解法 1：特殊处理头节点
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
