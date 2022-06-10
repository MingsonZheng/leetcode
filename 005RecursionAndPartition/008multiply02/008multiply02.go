package main

import (
	"fmt"
	"math"
)

// 面试题 08.05. 递归乘法
// 解法 2：min, max

func main() {
	fmt.Println(multiply(1, 10))
	fmt.Println(multiply(3, 4))
}

func multiply(A int, B int) int {
	if A == 1 {
		return B
	}
	if B == 1 {
		return A
	}
	min := int(math.Min(float64(A), float64(B)))
	max := int(math.Max(float64(A), float64(B)))
	// min 个 max 相加 = (min/2 个 max 相加) + (min/2 个 max 相加) + 0 (或 max)
	halfValue := multiply(min/2, max)
	if min%2 == 1 {
		return halfValue + halfValue + max
	} else {
		return halfValue + halfValue
	}
}
