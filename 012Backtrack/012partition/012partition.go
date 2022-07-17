package main

import "fmt"

// 131. 分割回文串

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

var result [][]string

func partition(s string) [][]string {
	result = make([][]string, 0)
	backtrack(s, 0, []string{})
	return result
}

func backtrack(s string, k int, path []string) {
	if k == len(s) {
		snapshot := make([]string, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	for end := k; end < len(s); end++ {
		if ispalidrome(s, k, end) {
			path = append(path, s[k:end+1])
			backtrack(s, end+1, path)
			path = path[:len(path)-1]
		}
	}
}

func ispalidrome(s string, p int, r int) bool {
	i := p
	j := r
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
