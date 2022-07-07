package main

import "fmt"

// 面试题 08.12. 八皇后

func main() {
	fmt.Println(solveNQueens(4))
}

var result [][]string

func solveNQueens(n int) [][]string {
	result = make([][]string, 0)
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	backtrack(0, board, n)
	return result
}

// row：阶段
// board：路径，记录已经做出的决策
// 可选列表：通过 board 推导出来，没有显示记录
func backtrack(row int, board [][]byte, n int) {
	// 结束条件，得到可行解
	if row == n {
		snapshot := make([]string, len(board))
		for i := 0; i < n; i++ {
			snapshot[i] = string(board[i])
		}
		result = append(result, snapshot)
		return
	}
	for col := 0; col < n; col++ { // 每一行都有 n 种做法
		if isOk(board, n, row, col) { // 可选列表
			board[row][col] = 'Q'      // 做选择，第 row 行的棋子放在 col 列
			backtrack(row+1, board, n) // 考察下一行
			board[row][col] = '.'      // 恢复选择
		}
	}
}

// 判断 row 行 column 列放置是否合适
func isOk(board [][]byte, n, row, col int) bool {
	// 检查是否有冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上对角线是否有冲突
	i := row - 1
	j := col + 1
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	// 检查左下对角线是否有冲突
	i = row - 1
	j = col - 1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}
