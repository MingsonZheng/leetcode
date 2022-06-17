package main

import "fmt"

// 69. x 的平方根

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 从 [1, x] 中查找最后一个平方小于等于 x 的数
	low := 1
	high := x/2 + 1
	for low <= high {
		mid := low + (high-low)/2
		var mid2 = int64(mid * mid)
		if mid2 == int64(x) {
			return mid
		} else if mid2 < int64(x) {
			var mid22 = int64((mid + 1) * (mid + 1))
			if mid22 <= int64(x) {
				low = mid + 1
			} else {
				return mid
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
