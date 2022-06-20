package main

import "fmt"

// 136. 只出现一次的数字

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))

	nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	hashTable := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		count := 1
		if c, ok := hashTable[nums[i]]; ok {
			count += c
		}
		hashTable[nums[i]] = count
	}
	for i := 0; i < len(nums); i++ {
		count := hashTable[nums[i]]
		if count == 1 {
			return nums[i]
		}
	}
	return -1
}
