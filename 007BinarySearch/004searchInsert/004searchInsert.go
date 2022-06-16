package main

import "fmt"

// 35. 搜索插入位置

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
	target = 2
	fmt.Println(searchInsert(nums, target))
	target = 7
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return len(nums)
}
