package main

import "fmt"

// 74. 搜索二维矩阵

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
	target = 13
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	low := 0
	high := m*n - 1
	var mid, midValue int
	for low <= high {
		mid = low + (high-low)/2
		midValue = matrix[mid/n][mid%n]
		if target == midValue {
			return true
		} else if target < midValue {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
