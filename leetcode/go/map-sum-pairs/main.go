package main

import (
	"fmt"
	"strconv"
)

type TrieNode struct {
	isWord bool
	children map[string]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func TrieConstructor() Trie {
	return Trie{&TrieNode{children: make(map[string]*TrieNode)}}
}

func (T *Trie) Delete() {
	T.root.children = make(map[string]*TrieNode)
}

func (T *Trie) Insert(word string)  {
	cur := T.root
	for _, c := range word {
		if _, ok := cur.children[strconv.Itoa(int(c))]; !ok {
			cur.children[strconv.Itoa(int(c))] = &TrieNode{children: make(map[string]*TrieNode)}
		}
		cur = cur.children[strconv.Itoa(int(c))]
	}
	cur.isWord = true
}

func (T *Trie) Search(word string) bool {
	cur := T.root
	for _, c := range word {
		if _, ok := cur.children[strconv.Itoa(int(c))]; !ok {
			return false
		}
		cur = cur.children[strconv.Itoa(int(c))]
	}
	if cur.isWord == false {
		return false
	}
	return true
}

func (T *Trie) StartsWith(prefix string) bool {
	cur := T.root
	for _, c := range prefix {
		if _, ok := cur.children[strconv.Itoa(int(c))]; !ok {
			return false
		}
		cur = cur.children[strconv.Itoa(int(c))]
	}
	return true
}

type MapSum struct {
	items map[string]int
	tr *Trie
}

func Constructor() MapSum {
	tr := TrieConstructor()
	return MapSum{items: make(map[string]int), tr: &tr}
}

func (this *MapSum) Insert(key string, val int)  {
	this.items[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	var sum int
	for key,val := range this.items {
		this.tr.Insert(key)
		if this.tr.StartsWith(prefix) {
			sum += val
		}
		this.tr.Delete()
	}
	return sum
}

func main() {

	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	fmt.Println(mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	fmt.Println(mapSum.Sum("ap"))
}
