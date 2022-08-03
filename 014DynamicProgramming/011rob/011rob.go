package main

import (
	"fmt"
	"math"
)

// 213. 打家劫舍 II

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	max1 := rob_dp(nums, 1, n-1)           // 第一个偷，范围排除最后一个
	max2 := nums[0] + rob_dp(nums, 2, n-2) // 第一个不偷，范围排除第一个
	return int(math.Max(float64(max1), float64(max2)))
}

func rob_dp(nums []int, p, r int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[p][0] = 0
	dp[p][1] = nums[p]
	for i := p + 1; i <= r; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1])))
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return int(math.Max(float64(dp[r][0]), float64(dp[r][1])))
}
