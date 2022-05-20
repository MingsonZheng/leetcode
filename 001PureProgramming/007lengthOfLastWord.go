package main

import "fmt"

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

// 58. 最后一个单词的长度
func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	if i < 0 {
		return 0
	}

	len := 0
	for i >= 0 && s[i] != ' ' {
		len++
		i--
	}
	return len
}
