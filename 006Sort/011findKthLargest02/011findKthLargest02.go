package main

import "fmt"

// 215. 数组中的第K个最大元素
// 解法 2：双指针

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	if k <= 0 || len(nums) < k {
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

// 从小到大
func partition(nums []int, p, q int) int {
	r := q
	i := p
	j := q - 1
	for i <= j && i <= q && j >= p {
		if nums[i] >= nums[r] {
			i++
			continue
		}
		if nums[j] < nums[r] {
			j--
			continue
		}
		if nums[i] < nums[r] && nums[j] >= nums[r] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
