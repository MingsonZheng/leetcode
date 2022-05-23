package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := spiralOrder(matrix)
	for _, v := range res {
		fmt.Printf("%v，", v)
	}
}

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, 0)
	left := 0
	right := n - 1
	top := 0
	bottom := m - 1
	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if top != bottom {
			for j := right - 1; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
		}
		if left != right {
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
