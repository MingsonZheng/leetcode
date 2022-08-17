package main

import (
	"fmt"
	"math"
)

// 53. 最大子数组和
// 解法 1：滑动窗口

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	n := len(nums)
	maxSum := math.MinInt32
	sum := 0
	for i := 0; i < n; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
