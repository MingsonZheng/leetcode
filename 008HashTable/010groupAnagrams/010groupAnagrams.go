package main

import (
	"fmt"
	"sort"
)

// 49. 字母异位词分组

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string, 0)
	for _, str := range strs {
		array := []byte(str)
		sort.Slice(array, func(i, j int) bool {
			return array[i] < array[j]
		})
		key := string(array)
		list := make([]string, 0)
		if _, ok := groupMap[key]; ok {
			list = groupMap[key]
		}
		list = append(list, str)
		groupMap[key] = list
	}
	// 没有现成的取 map 里 values 的方法，这里一个一个获取
	result := make([][]string, 0)
	for _, value := range groupMap {
		result = append(result, value)
	}
	return result
}
