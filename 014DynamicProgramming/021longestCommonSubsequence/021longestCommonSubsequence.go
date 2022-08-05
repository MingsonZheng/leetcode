package main

import "fmt"

// 1143. 最长公共子序列

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	t1 := []byte(text1)
	t2 := []byte(text2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = 0
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if t1[i-1] == t2[j-1] {
				dp[i][j] = max3(dp[i-1][j-1]+1, dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = max3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func max3(a, b, c int) int {
	max := a
	if max < b {
		max = b
	}
	if max < c {
		max = c
	}
	return max
}
