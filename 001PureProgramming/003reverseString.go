package main

import "fmt"

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

	s2 := []byte("hello")
	reverseString2(s2)
	fmt.Println(string(s2))

	s666 := []byte("hello")
	reverseString2(s666)
	fmt.Println(string(s666))
}

// 344. 反转字符串
func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func reverseString2(s []byte) {
	n := len(s)
	i := 0
	j := n - 1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseString666(s []byte) {
	var temp byte
	l := len(s)
	for i := 0; i < l/2; i++ {
		temp = s[i]
		s[i] = s[l-1-i]
		s[l-1-i] = temp
	}
}
