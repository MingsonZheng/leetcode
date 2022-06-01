package main

import "fmt"

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}

func calculate(s string) int {
	nums := make([]int, 0)
	ops := make([]byte, 0)
	i := 0
	n := len(s)
	for i < n {
		c := s[i]
		if c == ' ' { // 跳过空格
			i++
		} else if isDigit(c) { // 数字
			number := 0
			for i < n && isDigit(s[i]) {
				number = number*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, number)
		} else { // 运算符
			if len(ops) == 0 || prior(c, ops[len(ops)-1]) {
				ops = append(ops, c)
			} else {
				for len(ops) > 0 && !prior(c, ops[len(ops)-1]) {
					fetchAndCal(&nums, &ops) // 这里得传指针类型
				}
				ops = append(ops, c)
			}
			i++
		}
	}
	for len(ops) > 0 {
		fetchAndCal(&nums, &ops)
	}
	return nums[len(nums)-1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func prior(a, b byte) bool {
	if (a == '*' || a == '/') && (b == '+' || b == '-') {
		return true
	}
	return false
}

func fetchAndCal(nums *[]int, ops *[]byte) {
	number2 := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	number1 := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]
	ret := cal(op, number1, number2)
	*nums = append(*nums, ret)
}

func cal(op byte, number1, number2 int) int {
	switch op {
	case '+':
		return number1 + number2
	case '-':
		return number1 - number2
	case '*':
		return number1 * number2
	case '/':
		return number1 / number2
	}
	return -1
}
