package main

import (
	"container/heap"
	"fmt"
)

// 求中位数：295. 数据流的中位数

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}

type MedianFinder struct {
	minQueue *minHeap
	maxQueue *maxHeap
}

func Constructor() MedianFinder {
	minHeap := &minHeap{}
	maxHeap := &maxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{
		minQueue: minHeap,
		maxQueue: maxHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxQueue.Len() == 0 || num <= this.maxQueue.Peek().(int) {
		heap.Push(this.maxQueue, num)
	} else {
		heap.Push(this.minQueue, num)
	}
	for this.maxQueue.Len() < this.minQueue.Len() {
		e := heap.Pop(this.minQueue)
		heap.Push(this.maxQueue, e)
	}
	for this.minQueue.Len() < this.maxQueue.Len()-1 {
		e := heap.Pop(this.maxQueue)
		heap.Push(this.minQueue, e)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxQueue.Len() > this.minQueue.Len() {
		return float64(this.maxQueue.Peek().(int))
	} else {
		return (float64(this.maxQueue.Peek().(int)) + float64(this.minQueue.Peek().(int))) / 2
	}
}
