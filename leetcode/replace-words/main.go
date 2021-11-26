package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TrieNode struct {
	isWord bool
	children map[string]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{children: make(map[string]*TrieNode)}}
}

func (T *Trie) Insert(word string) {
	cur := T.root
	for _, c := range word {
		if _, ok := cur.children[strconv.Itoa(int(c))]; !ok {
			cur.children[strconv.Itoa(int(c))] = &TrieNode{children: make(map[string]*TrieNode)}
		}
		cur = cur.children[strconv.Itoa(int(c))]
	}
	cur.isWord = true
}

func (T *Trie) InsertIfNotExist(word string) {
	cur := T.root
	for _, c := range word {
		if _, ok := cur.children[strconv.Itoa(int(c))]; !ok {
			cur.children[strconv.Itoa(int(c))] = &TrieNode{children: make(map[string]*TrieNode)}
		}
		cur = cur.children[strconv.Itoa(int(c))]
		//insert if not exist
		if cur.isWord == true {
			return
		}
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

func (T *Trie) DeleteAndNew() {
	T.root = &TrieNode{children: make(map[string]*TrieNode)}
}

func replaceWords(dictionary []string, sentence string) string {

	Tr := Constructor()
	var shortDict []string
	sort.Strings(dictionary)
	for _, v := range dictionary {
		Tr.InsertIfNotExist(v)
	}
	for _, v := range dictionary {
		if Tr.Search(v) {
			shortDict = append(shortDict, v)
		}
	}
	Tr.DeleteAndNew()

	sArray := strings.Split(sentence, " ")
	for i, sentenceItem := range sArray {
		Tr.Insert(sentenceItem)
		for _, dictItem := range shortDict {
			if Tr.StartsWith(dictItem) {
				sArray[i] = dictItem
				break
			}
		}
		Tr.DeleteAndNew()
	}
	return strings.Join(sArray, " ")
}

func main() {

	dictionary1 := []string{"cat","bat","rat"}
	sentence1 := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary1, sentence1))

	dictionary2 := []string{"catt","cat","bat","rat"}
	sentence2 := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary2, sentence2))
}
