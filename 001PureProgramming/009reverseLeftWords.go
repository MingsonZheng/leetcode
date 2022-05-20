package main

import "fmt"

func main() {
	s := "abcdefg"
	n := 2
	fmt.Println(reverseLeftWords(s, n))
	fmt.Println(reverseLeftWords2(s, n))
	fmt.Println(reverseLeftWords3(s, n))
}

// 剑指 Offer 58 - II. 左旋转字符串
// 往左移动 n 位
func reverseLeftWords(s string, n int) string {
	str := []byte(s)
	for i := 0; i < n; i++ { // 移动 n 次，每次左移 1 位
		tmp := str[0]
		for j := 1; j < len(str); j++ {
			str[j-1] = str[j]
		}
		str[len(str)-1] = tmp
	}
	return string(str)
}

// tmp 数组
func reverseLeftWords2(s string, n int) string {
	tmp := make([]byte, len(s))
	// 数组分为 [0~n~len]，先把 0 ~ n-1 放在 tmp 后面
	for i := 0; i < n; i++ {
		tmp[i+len(s)-n] = s[i]
	}
	// 再把 n ~ len-1 放到 tmp 的前面
	for i := n; i < len(s); i++ {
		tmp[i-n] = s[i]
	}
	return string(tmp)
}

// 用 go slice 实现
func reverseLeftWords3(s string, n int) string {
	str := []byte(s)
	tmp := s[:n] // [0~n)，下标 0 ~ n-1，不含 n
	// str[n:] 表示 [n ~ len)，含 n
	str = append(str[n:], tmp...) // [0 ~ n) 后面
	return string(str)
}
