package main

import (
	"container/list"
	"fmt"
)

// 225. 用队列实现栈
// 解法 2：list 实现，push 倒腾；pop，peek 直接取

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
	n := this.queue.Len()
	this.queue.PushBack(x)
	for i := 0; i < n; i++ {
		ele := this.queue.Front()
		this.queue.Remove(ele)
		this.queue.PushBack(ele.Value.(int))
	}
}

func (this *MyStack) Pop() int {
	ele := this.queue.Front()
	this.queue.Remove(ele)
	return ele.Value.(int)
}

func (this *MyStack) Top() int {
	return this.queue.Front().Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}
