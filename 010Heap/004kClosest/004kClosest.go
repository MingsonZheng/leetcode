package main

import (
	"container/heap"
	"fmt"
)

// 973. 最接近原点的 K 个点

func main() {
	points := [][]int{
		{1, 3},
		{-2, -2},
	}
	fmt.Println(kClosest(points, 1))
}

func kClosest(points [][]int, k int) [][]int {
	pq := &maxHeap{}
	heap.Init(pq)
	for i := 0; i < k; i++ {
		heap.Push(pq, []int{points[i][0]*points[i][0] + points[i][1]*points[i][1], i})
	}
	n := len(points)
	for i := k; i < n; i++ {
		dist := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		if dist < pq.Peek().([]int)[0] {
			heap.Pop(pq)
			heap.Push(pq, []int{dist, i})
		}
	}
	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = points[heap.Pop(pq).([]int)[1]]
	}
	return ans
}
