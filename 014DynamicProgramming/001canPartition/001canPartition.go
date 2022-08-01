package main

import "fmt"

// 416. 分割等和子集

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	// 因为 sum 是总和的一半，所以只要能找到等于 sum 的状态，另一半也必然等于 sum，不符合的上面已排除
	sum /= 2
	dp := make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true
	if nums[0] <= sum {
		dp[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= sum; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][sum]
}
