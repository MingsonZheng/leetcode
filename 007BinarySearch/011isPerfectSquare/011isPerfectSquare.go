package main

import "fmt"

// 367. 有效的完全平方数

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {
	// [1, num] 之间选择平方等于 num 的数
	low := 1
	high := num
	for low <= high {
		mid := low + (high-low)/2
		var mid2 = int64(mid * mid)
		if mid2 == int64(num) {
			return true
		} else if mid2 > int64(num) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
