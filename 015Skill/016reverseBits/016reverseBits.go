package main

import "fmt"

// 面试题 05.03. 翻转数位

func main() {
	num := 1775
	fmt.Println(reverseBits(num))
}

func reverseBits(num int) int {
	if num == 0 {
		return 1
	}
	nums := make([]int, 32)
	for i := 0; i < 32; i++ {
		nums[i] = num & 1
		num >>= 1
	}
	leftOneCounts := make([]int, 32)
	count := 0
	for i := 0; i < 32; i++ {
		leftOneCounts[i] = count
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
	}
	rightOneCounts := make([]int, 32)
	count = 0
	for i := 31; i >= 0; i-- {
		rightOneCounts[i] = count
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
	}
	ret := 0
	for i := 0; i < 32; i++ {
		if ret < leftOneCounts[i]+rightOneCounts[i]+1 {
			ret = leftOneCounts[i] + rightOneCounts[i] + 1
		}
	}
	return ret
}
