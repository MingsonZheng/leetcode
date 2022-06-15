package main

import "fmt"

// 面试题 17.14. 最小K个数

func main() {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	k := 4
	fmt.Println(smallestK(arr, k))
}

var result []int
var count = 0

func smallestK(arr []int, k int) []int {
	if k == 0 && len(arr) < k {
		return []int{}
	}
	result = make([]int, k)
	// 重新初始化，不然 leetcode 里上一个测试用例的结果当作下一个测试用例的初始值，可能是 leetcode 的 bug
	count = 0
	quickSort(arr, 0, len(arr)-1, k)
	return result
}

func quickSort(nums []int, p, r, k int) {
	if p > r {
		return
	}
	q := partition(nums, p, r)
	if q-p+1 == k {
		for i := p; i <= q; i++ {
			result[count] = nums[i]
			count++
		}
	} else if q-p+1 < k {
		for i := p; i <= q; i++ {
			result[count] = nums[i]
			count++
		}
		quickSort(nums, q+1, r, k-(q-p+1))
	} else {
		quickSort(nums, p, q-1, k)
	}
}

func partition(nums []int, p, r int) int {
	i := p - 1
	j := p
	for j < r {
		if nums[j] < nums[r] {
			nums[j], nums[i+1] = nums[i+1], nums[j]
			i++
		}
		j++
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
