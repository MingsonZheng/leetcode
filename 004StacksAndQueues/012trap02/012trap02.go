package main

import (
	"fmt"
	"math"
)

// 42. 接雨水
// 解法 2：前缀后缀统计解法

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	// 前缀 max
	leftMax := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		leftMax[i] = int(math.Max(float64(max), float64(height[i])))
		max = leftMax[i]
	}
	// 后缀 max
	rightMax := make([]int, n)
	max = 0
	for i := n - 1; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(max), float64(height[i])))
		max = rightMax[i]
	}
	// 每个柱子上承载的雨水
	result := 0
	for i := 1; i < n-1; i++ {
		result += int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - height[i]
	}
	return result
}
