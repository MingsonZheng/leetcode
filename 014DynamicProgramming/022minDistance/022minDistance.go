package main

import (
	"fmt"
	"math"
)

// 72. 编辑距离

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	w1 := []byte(word1)
	w2 := []byte(word2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = min3(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min3(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[n][m]
}

func min3(n1, n2, n3 int) int {
	return int(math.Min(float64(n1), math.Min(float64(n2), float64(n3))))
}
