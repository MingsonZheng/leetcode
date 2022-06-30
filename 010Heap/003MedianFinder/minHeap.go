package main

import "math"

type priorityQueue []int

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type minHeap struct {
	priorityQueue
}

func (pq minHeap) Less(i, j int) bool {
	return pq.priorityQueue[i] < pq.priorityQueue[j]
}

func (pq minHeap) Peek() interface{} {
	if pq.Len() == 0 {
		return math.MinInt32
	}
	return pq.priorityQueue[0]
}

type maxHeap struct {
	priorityQueue
}

func (pq maxHeap) Less(i, j int) bool {
	return pq.priorityQueue[i] > pq.priorityQueue[j]
}

func (pq maxHeap) Peek() interface{} {
	if pq.Len() == 0 {
		return math.MaxInt32
	}
	return pq.priorityQueue[0]
}
