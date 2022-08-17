package main

import "fmt"

// 面试题 05.07. 配对交换

func main() {
	fmt.Println(exchangeBits(2))
}

func exchangeBits(num int) int {
	ret := 0
	for i := 0; i <= 30; i += 2 {
		a1 := num & (1 << i)
		b1 := num & (1 << (i + 1))
		if a1 != 0 {
			ret |= 1 << (i + 1)
		}
		if b1 != 0 {
			ret |= 1 << i
		}
	}
	return ret
}
