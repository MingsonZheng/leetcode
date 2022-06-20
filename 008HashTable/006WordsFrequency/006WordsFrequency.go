package main

import "fmt"

// 面试题 16.02. 单词频率

func main() {
	book := []string{"i", "have", "an", "apple", "he", "have", "a", "pen"}
	obj := Constructor(book)
	fmt.Println(obj.Get("you"))
	fmt.Println(obj.Get("have"))
	fmt.Println(obj.Get("an"))
	fmt.Println(obj.Get("apple"))
	fmt.Println(obj.Get("pen"))
}

type WordsFrequency struct {
	wordMap map[string]int
}

func Constructor(book []string) WordsFrequency {
	wordMap := make(map[string]int, 0)
	for _, word := range book {
		count := 1
		if c, ok := wordMap[word]; ok {
			count += c
		}
		wordMap[word] = count
	}
	return WordsFrequency{wordMap: wordMap}
}

func (this *WordsFrequency) Get(word string) int {
	if _, ok := this.wordMap[word]; !ok {
		return 0
	}
	return this.wordMap[word]
}
