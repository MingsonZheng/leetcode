package main

import "fmt"

// 20. 有效的括号

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else { // 右括号
			if len(stack) == 0 {
				return false
			}
			popC := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c == ')' && popC != '(' {
				return false
			}
			if c == ']' && popC != '[' {
				return false
			}
			if c == '}' && popC != '{' {
				return false
			}
		}
	}
	return len(stack) == 0
}
