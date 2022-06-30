package main

import "fmt"

// 208. 实现 Trie (前缀树) 解法 2：map

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
	children map[byte]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			data:     '/',
			children: map[byte]*TrieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := p.children[c]; !ok {
			p.children[c] = &TrieNode{data: c, children: map[byte]*TrieNode{}}
		}
		p = p.children[c]
	}
	p.isEnding = true
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := p.children[c]; !ok {
			return false
		}
		p = p.children[c]
	}
	return p.isEnding
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := p.children[c]; !ok {
			return false
		}
		p = p.children[c]
	}
	return true
}
