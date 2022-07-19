package main

import "fmt"

// 22. 括号生成

func main() {
	fmt.Println(generateParenthesis(3))
}

var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	path := make([]byte, 2*n)
	backtrack(n, 0, 0, 0, path)
	return result
}

func backtrack(n, leftUsed, rightUsed, k int, path []byte) {
	if k == 2*n {
		result = append(result, string(path))
		return
	}
	if leftUsed < n {
		path[k] = '('
		backtrack(n, leftUsed+1, rightUsed, k+1, path)
	}
	if leftUsed > rightUsed {
		path[k] = ')'
		backtrack(n, leftUsed, rightUsed+1, k+1, path)
	}
}
