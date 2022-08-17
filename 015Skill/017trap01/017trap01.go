package main

import (
	"fmt"
	"math"
)

// 42. 接雨水
// 解法 1：暴力解法

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	result := 0
	// 遍历每个柱子 n，查找它左边的最高柱子 lh，和右边的最高柱子 rh
	// 柱子上能承载的雨水 = min(lh, rh)-h
	for i := 1; i < n-1; i++ {
		lh := 0
		for j := 0; j < i; j++ { // 左侧最高 lh
			if height[j] > lh {
				lh = height[j]
			}
		}
		rh := 0
		for j := i + 1; j < n; j++ {
			if height[j] > rh {
				rh = height[j]
			}
		}
		carry := int(math.Min(float64(lh), float64(rh))) - height[i]
		if carry < 0 {
			carry = 0
		}
		result += carry
	}
	return result
}
