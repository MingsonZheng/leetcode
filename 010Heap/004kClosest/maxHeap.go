package main

import "math"

type maxHeap [][]int

func (pq maxHeap) Len() int {
	return len(pq)
}
func (pq maxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *maxHeap) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *maxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq maxHeap) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq maxHeap) Peek() interface{} {
	if pq.Len() == 0 {
		return []int{math.MaxInt32, -1}
	}
	return pq[0]
}
