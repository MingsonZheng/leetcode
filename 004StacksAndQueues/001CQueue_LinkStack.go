package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}

// 剑指 Offer 09. 用两个栈实现队列
// go 标准库没有提供栈和队列,可以用以下三种方式代替:
// 1.用 list.List 实现, list.List 是双向链表
// 2.基于数组或链表实现栈和队列
// 3.slice

// Node go 标准库没有实现栈，这里用的自己实现的链式栈
type Node struct {
	data int
	next *Node
}

type LinkStack struct {
	head *Node
}

func (stack *LinkStack) Push(value int) {
	newNode := &Node{
		data: value,
		next: nil,
	}
	newNode.next = stack.head
	stack.head = newNode
}

func (stack *LinkStack) Pop() int {
	if stack.head == nil {
		return -1
	}
	value := stack.head.data
	stack.head = stack.head.next
	return value
}
func (stack *LinkStack) Peak() int {
	if stack.head == nil {
		return -1
	}
	return stack.head.data
}

func (stack *LinkStack) IsEmpty() bool {
	return stack.head == nil
}

//// CQueue 用两个栈实现队列
//type CQueue struct {
//	stack    LinkStack
//	tmpStack LinkStack
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		stack: LinkStack{
//			head: nil,
//		},
//		tmpStack: LinkStack{
//			head: nil,
//		},
//	}
//}
//
//// AppendTail 解法 2 自己实现链式栈，入队-倒腾两个栈；出队-直接出栈
//// 对应的 slice 实现,参考解法 4
//func (this *CQueue) AppendTail(value int) {
//	for !this.stack.IsEmpty() {
//		this.tmpStack.Push(this.stack.Pop())
//	}
//	this.stack.Push(value)
//	for !this.tmpStack.IsEmpty() {
//		this.stack.Push(this.tmpStack.Pop())
//	}
//}
//
//func (this *CQueue) DeleteHead() int {
//	if this.stack.IsEmpty() {
//		return -1
//	}
//	return this.stack.Pop()
//}

// AppendTail 解法 1 自己实现链式栈，入队-直接入栈；出队-倒腾两个栈
// 对应的 slice 实现,参考解法 3
//func (this *CQueue) AppendTail(value int) {
//	this.stack.Push(value)
//}
//
//func (this *CQueue) DeleteHead() int {
//	if this.stack.IsEmpty() {
//		return -1
//	}
//	for !this.stack.IsEmpty() {
//		this.tmpStack.Push(this.stack.Pop())
//	}
//	result := this.tmpStack.Pop()
//	for !this.tmpStack.IsEmpty() {
//		this.stack.Push(this.tmpStack.Pop())
//	}
//	return result
//}
