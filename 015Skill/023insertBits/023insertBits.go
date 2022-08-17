package main

import "fmt"

// 面试题 05.01. 插入

func main() {
	fmt.Println(insertBits(1024, 19, 2, 6))
}

func insertBits(N int, M int, i int, j int) int {
	nBits := make([]int, 32)
	mBits := make([]int, 32)
	mask := 1
	for k := 0; k < 32; k++ {
		if N&mask != 0 {
			nBits[k] = 1
		}
		mask <<= 1
	}
	mask = 1
	for k := 0; k < 32; k++ {
		if M&mask != 0 {
			mBits[k] = 1
		}
		mask <<= 1
	}

	for k := i; k <= j; k++ {
		nBits[k] = mBits[k-i]
	}
	mask = 1
	ret := 0
	for k := 0; k < 32; k++ {
		ret += nBits[k] * mask
		mask <<= 1
	}
	return ret
}
