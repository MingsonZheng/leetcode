package main

import "fmt"

// 面试题 01.02. 判定是否互为字符重排

func main() {
	s1 := "abc"
	s2 := "bca"
	fmt.Println(CheckPermutation(s1, s2))

	s1 = "abc"
	s2 = "bad"
	fmt.Println(CheckPermutation(s1, s2))
}

func CheckPermutation(s1 string, s2 string) bool {
	s1ht := make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		count := 1
		if c2, ok := s1ht[c]; ok {
			count += c2
		}
		s1ht[c] = count
	}
	// s2 去跟 s1 匹配
	for i := 0; i < len(s2); i++ {
		c := s2[i]
		if _, ok := s1ht[c]; !ok {
			return false
		}
		count := s1ht[c]
		if count == 0 {
			return false
		}
		s1ht[c] = count - 1
	}
	// 检查 s1ht 是否为空
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if s1ht[c] != 0 {
			return false
		}
	}
	return true
}
