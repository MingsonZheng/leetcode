package main

import "fmt"

// 494. 目标和

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(nums, 3))
}

func findTargetSumWays(nums []int, target int) int {
	if target > 1000 || target < -1000 {
		return 0
	}
	n := len(nums)
	offset := 1000 // 偏移 1000，[-1000, 1000] -> [0, 2000]
	w := 2000      // 背包重量，[0, 2000]
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w+1)
	}
	dp[0][offset-nums[0]] += 1 // 为什么是 +=，而不是 =，因为有可能多次计算后回到等于 0 的状态，代表的不是同一个状态
	dp[0][offset+nums[0]] += 1

	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if j-nums[i] >= 0 && j-nums[i] <= w {
				dp[i][j] = dp[i-1][j-nums[i]]
			}
			if j+nums[i] >= 0 && j+nums[i] <= w {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
		}
	}
	return dp[n-1][target+1000]
}
