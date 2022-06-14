package main

import "fmt"

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}

func exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]%2 == 1 {
			i++
			continue
		}
		if nums[j]%2 == 0 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
