package main

import (
	"fmt"
	"math"
)

// 198. 打家劫舍

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1])))
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return int(math.Max(float64(dp[n-1][0]), float64(dp[n-1][1])))
}
