package main

import "fmt"

// 75. 颜色分类

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	p := 0
	q := len(nums) - 1
	for p < q {
		if nums[p] != 2 {
			p++
			continue
		}
		if nums[q] == 2 {
			q--
			continue
		}
		nums[p], nums[q] = nums[q], nums[p]
		p++
		q--
	}
	i := 0
	j := p
	if nums[j] == 2 {
		j--
	}
	for i < j {
		if nums[i] == 0 {
			i++
			continue
		}
		if nums[j] == 1 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
