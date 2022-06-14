package main

import (
	"fmt"
	"sort"
)

// 242. 有效的字母异位词

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	str1 := []byte(s)
	str2 := []byte(t)

	sort.Slice(str1, func(i, j int) bool {
		return str1[i] < str1[j]
	})
	sort.Slice(str2, func(i, j int) bool {
		return str2[i] < str2[j]
	})

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}
