package main

import "fmt"

// 剑指 Offer 03. 数组中重复的数字

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	// go 标准库没有 set，用 map 实现
	set := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		if set[nums[i]] {
			return nums[i]
		}
		set[nums[i]] = true
	}
	return -1
}
