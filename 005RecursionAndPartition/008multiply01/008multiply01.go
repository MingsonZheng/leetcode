package main

import "fmt"

// 面试题 08.05. 递归乘法
// 解法 1：halfValue

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
	halfValue := multiply(A/2, B)
	if A%2 == 1 {
		return halfValue + halfValue + B
	} else {
		return halfValue + halfValue
	}
}
