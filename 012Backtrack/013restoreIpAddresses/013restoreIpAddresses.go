package main

import (
	"fmt"
	"strings"
)

// 93. 复原 IP 地址

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

var result []string

func restoreIpAddresses(s string) []string {
	result = make([]string, 0)
	backtrack(s, 0, 0, []int{})
	return result
}

// k：字符串 s 第 k 位
func backtrack(s string, k, step int, path []int) {
	if k == len(s) && step == 4 {
		sb := strings.Builder{}
		for i := 0; i < 3; i++ {
			sb.Write([]byte(fmt.Sprintf("%d", path[i])))
			sb.WriteByte(byte('.'))
		}
		sb.Write([]byte(fmt.Sprintf("%d", path[3])))
		result = append(result, sb.String())
	}
	if step > 4 {
		return
	}
	if k == len(s) {
		return
	}
	val := 0
	// 1 位数，0 ~ 9，必然小于 255
	if k < len(s) {
		val = val*10 + int(s[k]-'0')
		path = append(path, val)
		backtrack(s, k+1, step+1, path)
		path = path[:len(path)-1]
	}
	if s[k] == '0' { // 前导 0 不行
		return
	}
	// 2 位数，0 ~ 99，必然小于 255
	if k+1 < len(s) {
		val = val*10 + int(s[k+1]-'0')
		path = append(path, val)
		backtrack(s, k+2, step+1, path)
		path = path[:len(path)-1]
	}
	// 3 位数，0 ~ 255，未必小于 255
	if k+2 < len(s) {
		val = val*10 + int(s[k+2]-'0')
		if val <= 255 {
			path = append(path, val)
			backtrack(s, k+3, step+1, path)
			path = path[:len(path)-1]
		}
	}
}
