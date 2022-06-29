package main

import (
	"container/heap"
	"fmt"
)

// 优先级队列：23. 合并K个升序链表
// go 优先级队列可以用 heap 实现.  标准库给了个 example
// https://github.com/golang/go/blob/master/src/container/heap/example_pq_test.go

func main() {
	array1 := []int{1, 4, 5}
	list1 := initial(array1)
	travel(list1)
	fmt.Println()

	array2 := []int{1, 3, 4}
	list2 := initial(array2)
	travel(list2)
	fmt.Println()

	array3 := []int{2, 6}
	list3 := initial(array3)
	travel(list3)
	fmt.Println()

	lists := []*ListNode{list1, list2, list3}
	head := mergeKLists(lists)
	travel(head)
	fmt.Println()
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	k := len(lists)
	minQ := &minHeap{}
	heap.Init(minQ)
	for i := 0; i < k; i++ {
		if lists[i] != nil {
			heap.Push(minQ, lists[i])
		}
	}
	head := &ListNode{}
	tail := head
	for minQ.Len() > 0 {
		curNode := heap.Pop(minQ).(*ListNode)
		tail.Next = curNode
		tail = tail.Next
		if curNode.Next != nil {
			heap.Push(minQ, curNode.Next)
		}
	}
	return head.Next
}
