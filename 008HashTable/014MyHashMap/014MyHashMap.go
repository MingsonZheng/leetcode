package main

import "fmt"

// 706. 设计哈希映射

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj.Get(2))
}

type Pair struct {
	key   int
	value int
}

type MyHashMap struct {
	slots [][]Pair
}

const SLOTS_COUNT = 3535

func Constructor() MyHashMap {
	return MyHashMap{slots: make([][]Pair, SLOTS_COUNT)}
}

func (this *MyHashMap) hash(key int) int {
	return key % SLOTS_COUNT
}

func (this *MyHashMap) Put(key int, value int) {
	slot := this.slots[this.hash(key)]
	if slot == nil {
		this.slots[this.hash(key)] = make([]Pair, 0)
		slot = this.slots[this.hash(key)]
	}
	for i := 0; i < len(slot); i++ {
		pair := slot[i]
		if pair.key == key {
			slot = append(slot[:i], slot[i+1:]...) // remove i
			break
		}
	}
	slot = append(slot, Pair{key, value})
	// append 会返回新的 slice，因此这里还需要再放回去
	this.slots[this.hash(key)] = slot
}

func (this *MyHashMap) Get(key int) int {
	slot := this.slots[this.hash(key)]
	if slot == nil {
		return -1
	}
	for i := 0; i < len(slot); i++ {
		pair := slot[i]
		if pair.key == key {
			return pair.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	slot := this.slots[this.hash(key)]
	if slot == nil {
		return
	}
	for i := 0; i < len(slot); i++ {
		pair := slot[i]
		if pair.key == key {
			slot = append(slot[:i], slot[i+1:]...) // remove i
			break
		}
	}
	this.slots[this.hash(key)] = slot
}
