package main

import (
	"container/list"
	"fmt"
)

// 225. 用队列实现栈
// 解法 1：list 实现，push 直接塞；pop，peek 倒腾

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	queue list.List
}

func Constructor() MyStack {
	return MyStack{
		queue: list.List{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
}

func (this *MyStack) Pop() int {
	n := this.queue.Len()
	for i := 0; i < n-1; i++ {
		ele := this.queue.Front()
		this.queue.Remove(ele)
		this.queue.PushBack(ele.Value.(int))
	}
	ele := this.queue.Front()
	this.queue.Remove(ele)
	return ele.Value.(int)
}

func (this *MyStack) Top() int {
	n := this.queue.Len()
	for i := 0; i < n-1; i++ {
		ele := this.queue.Front()
		this.queue.Remove(ele)
		this.queue.PushBack(ele.Value.(int))
	}
	ele := this.queue.Front()
	this.queue.Remove(ele)
	this.queue.PushBack(ele.Value.(int))
	return ele.Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}
