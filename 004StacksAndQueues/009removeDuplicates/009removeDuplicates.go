package main

// 1047. 删除字符串中的所有相邻重复项

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(s string) string {
	deque := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(deque) == 0 || deque[len(deque)-1] != c {
			deque = append(deque, c)
		} else {
			deque = deque[:len(deque)-1]
		}
	}
	var sb strings.Builder
	for len(deque) > 0 {
		pollFirst := deque[0]
		deque = deque[1:]
		sb.WriteByte(pollFirst)
	}
	return sb.String()
}
