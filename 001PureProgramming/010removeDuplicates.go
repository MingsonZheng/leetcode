package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	k := 0 // [0,k] 栈中元素
	for i := 1; i < n; i++ {
		if nums[i] != nums[k] { // 放入栈
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
