package main

import "fmt"

// 面试题 03.05. 栈排序
// 解法 2：类似插入排序，一直保持栈中元素从大到小有序（从栈底到栈顶）

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	obj.Pop()
	fmt.Println(obj.Peek())
}

type SortedStack struct {
	stack    []int
	tmpStack []int
}

func Constructor() SortedStack {
	return SortedStack{
		stack:    make([]int, 0),
		tmpStack: make([]int, 0),
	}
}

func (this *SortedStack) Push(val int) {
	for len(this.stack) > 0 && this.stack[len(this.stack)-1] < val {
		top := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.tmpStack = append(this.tmpStack, top)
	}
	this.stack = append(this.stack, val)
	for len(this.tmpStack) > 0 {
		top := this.tmpStack[len(this.tmpStack)-1]
		this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
		this.stack = append(this.stack, top)
	}
}

func (this *SortedStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *SortedStack) Peek() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.stack) == 0
}
