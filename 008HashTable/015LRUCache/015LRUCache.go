package main

import "fmt"

// 146. LRU 缓存

func main() {
	capacity := 2
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

type LRUCache struct {
	hashtable map[int]*DLinkedNode
	size      int
	capacity  int
	head      *DLinkedNode
	tail      *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:      0,
		capacity:  capacity,
		hashtable: make(map[int]*DLinkedNode, 0),
		head:      &DLinkedNode{key: -1, value: -1},
		tail:      &DLinkedNode{key: -1, value: -1},
	}
	lruCache.head.prev = nil
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	lruCache.tail.next = nil
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if this.size == 0 {
		return -1
	}
	node := this.hashtable[key]
	if node == nil {
		return -1
	}
	this.removeNode(node)
	this.addNodeAtHead(node)
	return node.value
}

func (this *LRUCache) Remove(key int) {
	node := this.hashtable[key]
	if node != nil {
		this.removeNode(node)
		delete(this.hashtable, key)
		this.size--
		return
	}
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) addNodeAtHead(node *DLinkedNode) {
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) Put(key int, value int) {
	node := this.hashtable[key]
	if node != nil {
		node.value = value
		this.removeNode(node)
		this.addNodeAtHead(node)
		return
	}
	if this.size == this.capacity {
		delete(this.hashtable, this.tail.prev.key)
		this.removeNode(this.tail.prev)
		this.size--
	}
	newNode := &DLinkedNode{key: key, value: value}
	this.addNodeAtHead(newNode)
	this.hashtable[key] = newNode
	this.size++
}
