package main

import "fmt"

// 面试题 03.01. 三合一

func main() {
	obj := Constructor(1)
	obj.Push(0, 1)
	obj.Push(0, 2)
	fmt.Println(obj.Pop(0))
	fmt.Println(obj.Pop(0))
	fmt.Println(obj.Pop(0))
	fmt.Println(obj.IsEmpty(0))
}

type TripleInOne struct {
	array []int
	n     int
	top   []int // 保存每个栈的栈顶下标
}

func Constructor(stackSize int) TripleInOne {
	tripleInOne := TripleInOne{
		array: make([]int, 3*stackSize),
		n:     3 * stackSize,
		top:   make([]int, 3),
	}
	tripleInOne.top[0] = -3
	tripleInOne.top[1] = -2
	tripleInOne.top[2] = -1
	return tripleInOne

}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.top[stackNum]+3 >= this.n {
		return
	}
	this.top[stackNum] += 3
	this.array[this.top[stackNum]] = value
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.top[stackNum] < 0 {
		return -1
	}
	ret := this.array[this.top[stackNum]]
	this.top[stackNum] -= 3
	return ret
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.top[stackNum] < 0 {
		return -1
	}
	return this.array[this.top[stackNum]]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.top[stackNum] < 0
}
