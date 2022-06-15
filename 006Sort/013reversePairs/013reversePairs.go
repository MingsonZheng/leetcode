package main

import "fmt"

// 剑指 Offer 51. 数组中的逆序对

func main() {
	nums := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(nums))
}

var reverseCount = 0

func reversePairs(nums []int) int {
	// 每次都重新初始化，不然 leetcode 里上一个测试用例的结果会覆盖下一个测试用例的 reverseCount
	reverseCount = 0
	mergeSort(nums, 0, len(nums)-1)
	return reverseCount
}

func mergeSort(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(nums, p, q)
	mergeSort(nums, q+1, r)
	merge(nums, p, q, r)
}

func merge(nums []int, p, q, r int) int {
	tmp := make([]int, r-p+1)
	i := p
	j := q + 1
	k := 0
	for i <= q && j <= r {
		if nums[j] < nums[i] {
			reverseCount += q - i + 1
			tmp[k] = nums[j]
			k++
			j++
		} else {
			tmp[k] = nums[i]
			k++
			i++
		}
	}
	for j <= r {
		tmp[k] = nums[j]
		k++
		j++
	}
	for i <= q {
		tmp[k] = nums[i]
		k++
		i++
	}
	for i := 0; i < r-p+1; i++ {
		nums[i+p] = tmp[i]
	}
	return reverseCount
}
