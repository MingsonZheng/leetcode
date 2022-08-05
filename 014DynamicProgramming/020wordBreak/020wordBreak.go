package main

import "fmt"

// 139. 单词拆分

func main() {
	s := "leetcode"
	wordDict := []string{
		"leet",
		"code",
	}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			len := len(word)
			startPosition := i - len
			if startPosition >= 0 && startsWith(s, word, startPosition) && dp[i-len] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func startsWith(s, word string, start int) bool {
	return s[start:start+len(word)] == word
}
