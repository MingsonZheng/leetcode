package main

import "fmt"

func main() {
	board := []string{"O X", " XO", "X O"}
	fmt.Println(tictactoe(board))
	board2 := []string{"OOX", "XXO", "OXO"}
	fmt.Println(tictactoe(board2))
	board3 := []string{"OOX", "XXO", "OX "}
	fmt.Println(tictactoe(board3))
	board4 := []string{"OX ", "OX ", "O  "}
	fmt.Println(tictactoe(board4))
	board5 := []string{"OX ", "OO ", "XXO"}
	fmt.Println(tictactoe(board5))
}

// 面试题 16.04. 井字游戏
func tictactoe(board []string) string {
	n := len(board)
	boards := make([][]byte, n)

	for i := 0; i < len(boards); i++ {
		boards[i] = make([]byte, n)
		boards[i] = []byte(board[i])
	}

	determined := false // 表示是否已经有人赢了
	// 检查行
	for i := 0; i < n; i++ {
		if boards[i][0] == ' ' {
			continue
		}
		determined = true
		for j := 1; j < n; j++ {
			if boards[i][j] != boards[i][0] {
				determined = false
				break
			}
		}
		if determined {
			return string(boards[i][0])
		}
	}
	// 检查列
	for j := 0; j < n; j++ {
		if boards[0][j] == ' ' {
			continue
		}
		determined = true
		for i := 1; i < n; i++ {
			if boards[i][j] != boards[0][j] {
				determined = false
			}
		}
		if determined {
			return string(boards[0][j])
		}
	}
	// 检查对角线：左上->右下
	if boards[0][0] != ' ' {
		i := 1
		j := 1
		determined = true
		for i < n && j < n {
			if boards[i][j] != boards[0][0] {
				determined = false
				break
			}
			i++
			j++
		}
		if determined {
			return string(boards[0][0])
		}
	}
	// 检查对角线：左下->右上
	if boards[n-1][0] != ' ' {
		i := n - 2
		j := 1
		determined = true
		for i >= 0 && j < n {
			if boards[i][j] != boards[n-1][0] {
				determined = false
				break
			}
			i--
			j++
		}
		if determined {
			return string(boards[n-1][0])
		}
	}
	// 上面没有找到哪方赢，判断游戏是否还能继续玩
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if boards[i][j] == ' ' {
				return "Pending"
			}
		}
	}
	// 游戏结束了，平局
	return "Draw"
}
