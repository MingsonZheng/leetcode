package main

import "fmt"

// 438. 找到字符串中所有字母异位词

func main() {
	//s := "cbaebabacd"
	//p := "abc"
	//fmt.Println(findAnagrams(s, p))
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	n := len(s)
	m := len(p)
	if m > n {
		return []int{}
	}
	needs := make([]int, 26)
	for i := 0; i < m; i++ {
		needs[p[i]-'a']++
	}
	matched := make([]int, 26)
	startP := 0
	endP := 0
	result := make([]int, 0)
	for endP < m {
		matched[s[endP]-'a']++
		endP++
	}
	if same(needs, matched) {
		result = append(result, startP)
	}
	for endP < n && startP < n {
		matched[s[startP]-'a']--
		matched[s[endP]-'a']++
		startP++
		endP++
		if same(needs, matched) {
			result = append(result, startP)
		}
	}
	return result
}

func same(needs, matched []int) bool {
	for i := 0; i < len(needs); i++ {
		if needs[i] != matched[i] {
			return false
		}
	}
	return true
}
