package main

import "fmt"

// 153. 寻找旋转排序数组中的最小值

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums))
	nums = []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		// 特殊处理 low == high 的情况
		if low == high {
			return nums[mid]
		}
		// 先处理命中情况
		if (mid != 0 && nums[mid] < nums[mid-1]) || (mid == 0 && nums[mid] < nums[high]) {
			return nums[mid]
		} else if nums[mid] > nums[high] { // 右循环有序
			low = mid + 1
		} else { // 右侧非循环有序
			high = mid - 1
		}
	}
	return -1 // 永远到不了这里
}
