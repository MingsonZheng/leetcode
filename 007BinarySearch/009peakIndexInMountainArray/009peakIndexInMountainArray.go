package main

import "fmt"

// 852. 山脉数组的峰顶索引

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
	arr = []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
	arr = []int{0, 10, 5, 2}
	fmt.Println(peakIndexInMountainArray(arr))
	arr = []int{3, 4, 5, 1}
	fmt.Println(peakIndexInMountainArray(arr))
	arr = []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if mid == 0 {
			low = mid + 1
		} else if mid == n-1 {
			high = mid - 1
		} else if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
