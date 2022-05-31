package main

import (
	"fmt"
	"math"
)

// 面试题 03.05. 栈排序
// 解法 1：pop，peek 倒腾

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
	this.stack = append(this.stack, val)
}

func (this *SortedStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	minVal := math.MaxInt32
	for len(this.stack) > 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		if val < minVal {
			minVal = val
		}
		this.tmpStack = append(this.tmpStack, val)
	}
	removed := false // 标记是否已经 remove 了，如果有多个最小值，只 remove 一个
	for len(this.tmpStack) > 0 {
		val := this.tmpStack[len(this.tmpStack)-1]
		this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
		if (val != minVal) || (val == minVal && removed == true) {
			this.stack = append(this.stack, val)
		} else {
			removed = true
		}
	}
}

func (this *SortedStack) Peek() int {
	if len(this.stack) == 0 {
		return -1
	}
	minVal := math.MaxInt32
	for len(this.stack) > 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		if val < minVal {
			minVal = val
		}
		this.tmpStack = append(this.tmpStack, val)
	}
	for len(this.tmpStack) > 0 {
		val := this.tmpStack[len(this.tmpStack)-1]
		this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
		this.stack = append(this.stack, val)
	}
	return minVal
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.stack) == 0
}
