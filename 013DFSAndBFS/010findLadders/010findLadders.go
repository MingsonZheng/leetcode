package main

import "fmt"

// 面试题 17.22. 单词转换

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

var visited map[string]bool
var resultPath []string
var found bool

func findLadders(beginWord string, endWord string, wordList []string) []string {
	visited = make(map[string]bool, 0)
	resultPath = make([]string, 0)
	found = false
	dfs(beginWord, endWord, []string{}, wordList)
	return resultPath
}
func dfs(curWord, endWord string, path []string, wordlist []string) {
	if found {
		return
	}
	path = append(path, curWord)
	visited[curWord] = true
	if curWord == endWord {
		resultPath = append(resultPath, path...)
		found = true
		return
	}
	for i := 0; i < len(wordlist); i++ {
		nextWord := wordlist[i]
		if visited[nextWord] || !isValidChange(curWord, nextWord) {
			continue
		}
		dfs(nextWord, endWord, path, wordlist)
	}
	path = path[:len(path)-1]
}

func isValidChange(word1, word2 string) bool {
	diff := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
		}
	}
	return diff == 1
}
