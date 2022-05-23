package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate2(matrix)
	for _, array := range matrix {
		for _, v := range array {
			fmt.Printf("%v，", v)
		}
		fmt.Println()
	}
}

// 48. 旋转图像
// 解法 1 借助辅助数组
func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[j][n-i-1] = matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = tmp[i][j]
		}
	}
}

// 解法 2 用翻转代替旋转
func rotate2(matrix [][]int) {
	n := len(matrix)
	// 先上下翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			swap(matrix, i, j, n-i-1, j)
		}
	}
	// 再对角翻转（左上 - 右下）
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			swap(matrix, i, j, j, i)
		}
	}
}

func swap(matrix [][]int, i, j, p, q int) {
	matrix[i][j], matrix[p][q] = matrix[p][q], matrix[i][j]
}
