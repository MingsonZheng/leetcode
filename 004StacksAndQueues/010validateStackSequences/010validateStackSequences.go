package main

import "fmt"

// 剑指 Offer 31. 栈的压入、弹出序列

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
	popped = []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	j := 0
	for i := 0; i < len(popped); i++ {
		number := popped[i]
		if len(stack) > 0 && stack[len(stack)-1] == number {
			stack = stack[:len(stack)-1]
		} else {
			for j < len(pushed) && pushed[j] != number {
				stack = append(stack, pushed[j])
				j++
			}
			if j == len(pushed) {
				return false
			}
			j++
		}
	}
	return true
}
