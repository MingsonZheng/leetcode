package main

type minHeap []*ListNode

func (pq minHeap) Len() int {
	return len(pq)
}
func (pq minHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// 思考了下为什么 Push 和 Pop 需要 *minHeap，其他的可以是 minHeap
// 1 是 minHeap 实现的是 heap 中的接口
// 2 是如果不用指针，minHeap 是 slice，append 会返回新的 pq，但是新的 pq 只是 push 方法的局部变量，所以 pq 并不会改变

func (pq *minHeap) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

// Pop 也是因为如果不用指针改变的只是局部变量
func (pq *minHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq minHeap) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
