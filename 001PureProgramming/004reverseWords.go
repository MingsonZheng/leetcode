package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world!  "
	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords666(s))
}

// 剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	str := []byte(s)

	n := trim(str)
	if n == 0 {
		return ""
	}
	reverse(str, 0, n-1)

	for p := 0; p < n; {
		r := p
		for r < n && str[r] != ' ' {
			r++
		}
		reverse(str, p, r-1)
		p = r + 1
	}
	// 这里只是为了配合输出
	newStr := make([]byte, n)
	for i := 0; i < n; i++ {
		newStr[i] = str[i]
	}
	return string(newStr)
}

// 原地删除前置空格和后置空格，以及内部多余空格，返回新字符串长度，单词之间只留一个空格
func trim(str []byte) int {
	i := 0
	n := len(str)
	k := 0 // 记录删除多余空格之后的数组长度
	for i < n && str[i] == ' ' {
		i++
	}
	for i < n {
		if str[i] == ' ' { // 删除内部多余的空格和末尾空格
			if i+1 < n && str[i+1] != ' ' {
				str[k] = ' '
				k++
			}
		} else {
			str[k] = str[i]
			k++
		}
		i++
	}
	return k
}

// 返回 [p,r] 之间的字符串，注意这里是闭区间
// 当然，前开后闭区间也可以，但代码中 i <= mid 应该改为 i < mid
func reverse(str []byte, p, r int) {
	mid := (p + r) / 2
	for i := p; i <= mid; i++ {
		str[i], str[r-(i-p)] = str[r-(i-p)], str[i]
	}
}

func reverseWords666(s string) string {
	res := ""
	first := true
	words := strings.Split(s, " ")
	n := len(words)
	for i := n - 1; i >= 0; i-- {
		// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括
		if words[i] != "" {
			// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个
			if first == false {
				res += " "
			}
			first = false
			res += words[i]
		}
	}
	return res
}
