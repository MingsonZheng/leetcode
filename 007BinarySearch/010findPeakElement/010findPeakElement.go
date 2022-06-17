package main

import "fmt"

// 162. 寻找峰值

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
	nums = []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	n := len(nums)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if low == high {
			return mid
		}
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				low = mid + 1
			}
		} else if mid == n-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			} else {
				high = mid - 1
			}
		} else if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
