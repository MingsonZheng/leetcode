package main

import "fmt"

// 121. 买卖股票的最佳时机

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	n := len(prices)
	max := make([]int, n)
	curMax := 0
	for i := n - 1; i >= 0; i-- {
		max[i] = curMax
		if prices[i] > curMax {
			curMax = prices[i]
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		if result < max[i]-prices[i] {
			result = max[i] - prices[i]
		}
	}
	return result
}
