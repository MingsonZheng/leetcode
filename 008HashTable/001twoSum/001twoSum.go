package main

import "fmt"

// 1. 两数之和

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	// 哈希表，key 是数本身，value 是下标
	hashTable := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		hashTable[nums[i]] = i
	}
	for i := 0; i < n; i++ {
		if value, ok := hashTable[target-nums[i]]; ok {
			if value != i {
				return []int{i, value}
			}
		}
	}
	return []int{}
}
