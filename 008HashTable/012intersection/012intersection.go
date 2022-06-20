package main

import "fmt"

// 349. 两个数组的交集

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	hashTable := make(map[int]bool, 0)
	for i := 0; i < len(nums1); i++ {
		hashTable[nums1[i]] = true
	}
	result := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if hashTable[nums2[i]] {
			// delete 或者 hashTable[nums2[i]] = false
			delete(hashTable, nums2[i])
			result = append(result, nums2[i])
		}
	}
	return result
}
