package main

import "fmt"

// 17. 电话号码的字母组合

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	digits = ""
	fmt.Println(letterCombinations(digits))
	digits = "2"
	fmt.Println(letterCombinations(digits))
}

var result []string

func letterCombinations(digits string) []string {
	result = make([]string, 0)
	if len(digits) == 0 {
		return []string{}
	}
	mappings := make([]string, 10)
	mappings[2] = "abc"
	mappings[3] = "def"
	mappings[4] = "ghi"
	mappings[5] = "jkl"
	mappings[6] = "mno"
	mappings[7] = "pqrs"
	mappings[8] = "tuv"
	mappings[9] = "wxyz"
	path := make([]byte, len(digits))
	backtrack(mappings, digits, 0, path)
	return result
}

func backtrack(mappings []string, digits string, k int, path []byte) {
	if k == len(digits) {
		result = append(result, string(path))
		return
	}
	mapping := mappings[digits[k]-'0']
	for i := 0; i < len(mapping); i++ {
		path[k] = mapping[i]
		backtrack(mappings, digits, k+1, path)
	}
}
