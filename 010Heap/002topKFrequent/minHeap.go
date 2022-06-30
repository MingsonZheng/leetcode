package main

import "math"

type QElement struct {
	val   int
	count int
}

type minHeap []QElement

func (pq minHeap) Len() int {
	return len(pq)
}
func (pq minHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *minHeap) Push(x interface{}) {
	*pq = append(*pq, x.(QElement))
}

func (pq *minHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq minHeap) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

func (pq minHeap) Peek() interface{} {
	if pq.Len() == 0 {
		return QElement{math.MinInt32, -1}
	}
	return pq[0]
}
