package main

import "fmt"

// 215. 数组中的第K个最大元素

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	return quickSort(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, p, r, k int) int {
	if p > r {
		return -1
	}
	q := partition(nums, p, r)
	if q-p+1 == k {
		return nums[q]
	} else if q-p+1 < k {
		return quickSort(nums, q+1, r, k-(q-p+1))
	} else {
		return quickSort(nums, p, q-1, k)
	}
}

func partition(nums []int, p, r int) int {
	i := p - 1
	j := p
	for j < r {
		if nums[j] > nums[r] {
			nums[j], nums[i+1] = nums[i+1], nums[j]
			i++
		}
		j++
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
