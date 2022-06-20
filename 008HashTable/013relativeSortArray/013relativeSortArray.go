package main

import (
	"fmt"
	"sort"
)

// 1122. 数组的相对排序

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))

	arr1 = []int{28, 6, 22, 8, 44, 17}
	arr2 = []int{22, 28, 8, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// arr2 中每个数字在 arr1 中出现的次数
	counts := make(map[int]int, 0)
	// 先用 arr2 构建 hash 表
	for i := 0; i < len(arr2); i++ {
		counts[arr2[i]] = 0
	}
	// 扫描 arr1 统计 arr2 中每个数字在 arr1 中出现的次数
	for i := 0; i < len(arr1); i++ {
		if oldCount, ok := counts[arr1[i]]; ok {
			counts[arr1[i]] = oldCount + 1
		}
	}
	result := make([]int, len(arr1))
	k := 0
	// 将 counts 的数据按照 arr2 的顺序输出
	for i := 0; i < len(arr2); i++ {
		count := counts[arr2[i]]
		for j := 0; j < count; j++ {
			result[k+j] = arr2[i]
		}
		k += count
	}
	// 将 arr1 中未出现在 arr2 中的数据有序输出到 result
	sort.Ints(arr1)
	for i := 0; i < len(arr1); i++ {
		if _, ok := counts[arr1[i]]; !ok {
			result[k] = arr1[i]
			k++
		}
	}
	return result
}
