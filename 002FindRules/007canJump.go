package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

// 55. 跳跃游戏
func canJump(nums []int) bool {
	reachedMax := 0
	for i := 0; i < len(nums); i++ {
		// 表示前边任意一个走法都到达不了！
		if i > reachedMax {
			return false
		}
		if i+nums[i] > reachedMax {
			reachedMax = i + nums[i]
		}
		if reachedMax >= len(nums)-1 {
			return true
		}
	}
	return false
}
