package main

import "fmt"

// 208. 实现 Trie (前缀树) 解法 1：数组

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}

type TrieNode struct {
	data     byte
	isEnding bool
	children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			data:     '/',
			children: [26]*TrieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if p.children[c-'a'] == nil {
			p.children[c-'a'] = &TrieNode{data: c}
		}
		p = p.children[c-'a']
	}
	p.isEnding = true
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if p.children[c-'a'] == nil {
			return false
		}
		p = p.children[c-'a']
	}
	return p.isEnding
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if p.children[c-'a'] == nil {
			return false
		}
		p = p.children[c-'a']
	}
	return true
}
