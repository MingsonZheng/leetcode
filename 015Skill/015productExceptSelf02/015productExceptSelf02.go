package main

import "fmt"

// 238. 除自身以外数组的乘积

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return result
}
