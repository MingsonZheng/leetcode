package main

import "fmt"

// 34. 在排序数组中查找元素的第一个和最后一个位置

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
	target = 6
	fmt.Println(searchRange(nums, target))
	nums = []int{}
	target = 0
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	low := 0
	high := len(nums) - 1
	left := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				left = mid
				break
			} else {
				high = mid - 1
			}
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	low = 0
	high = len(nums) - 1
	right := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				right = mid
				break
			} else {
				low = mid + 1
			}
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return []int{left, right}
}
