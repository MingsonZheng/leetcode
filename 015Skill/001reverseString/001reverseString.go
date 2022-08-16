package main

import "fmt"

// 344. 反转字符串

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	n := len(s)
	i := 0
	j := n - 1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
