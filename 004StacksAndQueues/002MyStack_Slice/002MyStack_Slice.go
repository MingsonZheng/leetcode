package main

import (
	"fmt"
)

// 225. 用队列实现栈
// 解法 3：slice 实现，push 倒腾；pop，peek 直接取

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	n := len(this.queue)
	this.queue = append(this.queue, x)
	for i := 0; i < n; i++ {
		ele := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, ele)
	}
}

func (this *MyStack) Pop() int {
	ele := this.queue[0]
	this.queue = this.queue[1:]
	return ele
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
