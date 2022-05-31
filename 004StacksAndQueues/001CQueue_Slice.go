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

type CQueue struct {
	stack    []int
	tmpStack []int
}

func Constructor() CQueue {
	return CQueue{
		stack:    make([]int, 0),
		tmpStack: make([]int, 0),
	}
}

// AppendTail 解法 4，用 slice 实现，入队-倒腾两个栈；出队-直接出栈
func (this *CQueue) AppendTail(value int) {
	for len(this.stack) > 0 {
		//result := this.stack[len(this.stack)-1]// 表示最右边元素，这里表示取出，并未出栈
		this.tmpStack = append(this.tmpStack, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1] // 移除 this.stack 最右边元素，即出栈
	}
	this.stack = append(this.stack, value) // 追加到数组最右边，即入栈
	for len(this.tmpStack) > 0 {
		this.stack = append(this.stack, this.tmpStack[len(this.tmpStack)-1])
		this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
	}
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack) == 0 {
		return -1
	}
	result := this.stack[len(this.stack)-1] // this.stack 最右边元素
	this.stack = this.stack[:len(this.stack)-1]
	return result
}

//// AppendTail 解法 3，用 slice 实现，入队-直接入栈；出队-倒腾两个栈
//func (this *CQueue) AppendTail(value int) {
//	this.stack = append(this.stack, value) // 追加到数组最右边，即入栈
//}
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.stack) == 0 {
//		return -1
//	}
//	for len(this.stack) > 0 {
//		//result := this.stack[len(this.stack)-1]// 表示最右边元素，这里表示取出，并未出栈
//		this.tmpStack = append(this.tmpStack, this.stack[len(this.stack)-1])
//		this.stack = this.stack[:len(this.stack)-1] // 移除 this.stack 最右边元素，即出栈
//	}
//	result := this.tmpStack[len(this.tmpStack)-1] // this.tmpStack 最右边元素
//	this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
//	for len(this.tmpStack) > 0 {
//		this.stack = append(this.stack, this.tmpStack[len(this.tmpStack)-1])
//		this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
//	}
//	return result
//}
