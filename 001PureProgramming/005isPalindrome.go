package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

// 125. 验证回文串
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		// 不是数字或字母的话，i 就一直 ++
		if !isAlpha(s[i]) {
			i++
			continue
		}
		// 不是数字或字母的话，j 就一直 --
		if !isAlpha(s[j]) {
			j--
			continue
		}
		// 走到这里的话，i，j 都指向数字或字母，看看两个字符是否相符
		if toLower(s[i]) != toLower(s[j]) {
			return false
		} else {
			// i 和 j 往中间挪一位
			i++
			j--
		}
	}
	return true
}

// 大写转小写
func toLower(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c
	}
	if c >= '0' && c <= '9' {
		return c
	}
	// ASCII 码：大写 A ~ Z 65 ~ 90，小写 a ~ z 97 ~ 122
	return c + 32
}

// 判断是不是数字或字母
func isAlpha(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
