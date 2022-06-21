package main

import "fmt"

// 面试题 16.21. 交换和

func main() {
	array1 := []int{4, 1, 2, 1, 1, 2}
	array2 := []int{3, 6, 3, 3}
	fmt.Println(findSwapValues(array1, array2))

	array1 = []int{1, 2, 3}
	array2 = []int{4, 5, 6}
	fmt.Println(findSwapValues(array1, array2))
}

func findSwapValues(array1 []int, array2 []int) []int {
	n := len(array1)
	m := len(array2)
	// 计算数组 1 的和
	sum1 := 0
	for i := 0; i < n; i++ {
		sum1 += array1[i]
	}
	// 计算数组 2 的和，并且将元素放到哈希表中
	sum2 := 0
	hashTable := make(map[int]bool, 0)
	for i := 0; i < m; i++ {
		sum2 += array2[i]
		hashTable[array2[i]] = true
	}
	// sum1 + sum2 是奇数，无解
	sum := sum1 + sum2
	if sum%2 == 1 {
		return []int{}
	}
	// 遍历数组 1 中的每个元素，在哈希表中查找
	diff := sum/2 - sum1
	for i := 0; i < n; i++ {
		target := array1[i] + diff
		if hashTable[target] {
			return []int{array1[i], target}
		}
	}
	return []int{}
}
