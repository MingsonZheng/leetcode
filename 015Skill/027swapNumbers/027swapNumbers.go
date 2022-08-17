package main

import "fmt"

// 面试题 16.01. 交换数字

func main() {
	numbers := []int{1, 2}
	fmt.Println(swapNumbers(numbers))
}

func swapNumbers(numbers []int) []int {
	if numbers[0] == numbers[1] {
		return numbers
	}
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}
