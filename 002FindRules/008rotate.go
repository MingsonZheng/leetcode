package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate3(matrix)
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

// 解法 3 标准原地旋转
func rotate3(matrix [][]int) {
	n := len(matrix) // 3
	s1_i := 0
	s1_j := 0
	for n > 1 { // 四个角落的坐标
		s2_i := s1_i         // 0
		s2_j := s1_j + n - 1 // 2
		s3_i := s1_i + n - 1 // 2
		s3_j := s1_j + n - 1 // 2
		s4_i := s1_i + n - 1 // 2
		s4_j := s1_j         // 0

		for move := 0; move <= n-2; move++ {
			p1_i := s1_i + 0
			p1_j := s1_j + move
			p2_i := s2_i + move
			p2_j := s2_j + 0
			p3_i := s3_i + 0
			p3_j := s3_j - move
			p4_i := s4_i - move
			p4_j := s4_j + 0
			swap2(matrix, p1_i, p1_j, p2_i, p2_j, p3_i, p3_j, p4_i, p4_j)
		}
		s1_i++
		s1_j++
		n -= 2
	}
}

func swap2(a [][]int, i1, j1, i2, j2, i3, j3, i4, j4 int) {
	tmp := a[i1][j1]
	a[i1][j1] = a[i4][j4]
	a[i4][j4] = a[i3][j3]
	a[i3][j3] = a[i2][j2]
	a[i2][j2] = tmp
}
