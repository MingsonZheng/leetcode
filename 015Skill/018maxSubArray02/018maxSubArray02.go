package main

import (
	"fmt"
	"math"
)

// 53. 最大子数组和
// 解法 2：前后缀

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := make([]int, len(nums))
	max := make([]int, len(nums))
	curSum := 0
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		sum[i] = curSum
	}
	curMax := math.MinInt32
	for i := len(sum) - 1; i >= 0; i-- {
		if curMax < sum[i] {
			curMax = sum[i]
		}
		max[i] = curMax
	}
	result := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if i == 0 && result < max[0] {
			result = max[0]
		}
		if i != 0 && result < max[i]-sum[i-1] {
			result = max[i] - sum[i-1]
		}
	}
	return result
}
