package main

import "fmt"

// 300. 最长递增子序列

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

// 解法二
func lengthOfLIS(nums []int) int {
	n := len(nums)
	listToMin := make([]int, n+1)
	k := 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		len := binarySearch(listToMin, k, nums[i])
		if len == -1 {
			dp[i] = 1
		} else {
			dp[i] = len + 1
		}
		if dp[i] > k {
			k = dp[i]
			listToMin[dp[i]] = nums[i]
		} else if listToMin[dp[i]] > nums[i] {
			listToMin[dp[i]] = nums[i]
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

// binarySearch 查找最后一个比 target 小的元素位置
func binarySearch(a []int, k, target int) int {
	low := 1
	high := k
	for low <= high {
		mid := (low + high) / 2
		if a[mid] < target {
			if mid == k || a[mid+1] >= target {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
