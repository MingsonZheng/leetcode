package main

import "fmt"

// 面试题 05.06. 整数转换

func main() {
	fmt.Println(convertInteger(29, 15))
}

func convertInteger(A int, B int) int {
	C := A ^ B
	count := 0
	for i := 0; i < 32; i++ {
		if (C & 1) == 1 {
			count++
		}
		C >>= 1
	}
	return count
}
