package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

// 剑指 Offer 05. 替换空格
// 和 1108. IP 地址无效化 相同
func replaceSpace(s string) string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			sb.WriteByte(c)
		} else {
			sb.Write([]byte("%20"))
		}
	}
	return sb.String()
}
