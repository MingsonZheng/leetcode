package main

import "fmt"

// 155. 最小栈

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

type MinStack struct {
	data   []int
	minval []int
}

func Constructor() MinStack {
	return MinStack{
		data:   make([]int, 0),
		minval: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.data = append(this.data, val)
		this.minval = append(this.minval, val)
	} else {
		curminval := this.minval[len(this.minval)-1]
		if val < curminval {
			this.minval = append(this.minval, val)
		} else {
			this.minval = append(this.minval, curminval)
		}
		this.data = append(this.data, val)
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minval = this.minval[:len(this.minval)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minval[len(this.minval)-1]
}
