package main

import (
	"container/heap"
	"fmt"
)

// TOP K：347. 前 K 个高频元素

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, 0)
	for _, num := range nums {
		count := 0
		if c, ok := counts[num]; ok {
			count = c
		}
		counts[num] = count + 1
	}
	queue := &minHeap{}
	for num, count := range counts {
		if queue.Len() < k {
			heap.Push(queue, QElement{val: num, count: count})
		} else {
			if queue.Peek().(QElement).count < count {
				heap.Pop(queue)
				heap.Push(queue, QElement{val: num, count: count})
			}
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(queue).(QElement).val
	}
	return result
}
