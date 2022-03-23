package main

import "fmt"

type DictNode struct {
	isWord   bool
	children map[string]*DictNode
}

type WordDictionary struct {
	root *DictNode
}

func Constructor() WordDictionary {
	return WordDictionary{&DictNode{children: make(map[string]*DictNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range word {
		if _, ok := cur.children[string(c)]; !ok {
			cur.children[string(c)] = &DictNode{children: make(map[string]*DictNode)}
		}
		cur = cur.children[string(c)]
	}
	cur.isWord = true
}

func (this *WordDictionary) SearchFrom(from *DictNode, word string) bool {
	cur := from
	for i, c := range word {
		if string(c) == "." {
			for _, val := range cur.children {
				if this.SearchFrom(val, word[i+1:]) {
					return true
				}
			}
			return false
		} else if _, ok := cur.children[string(c)]; !ok {
			return false
		}
		cur = cur.children[string(c)]
	}
	if cur.isWord {
		return true
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	cur := this.root
	for i, c := range word {
		if string(c) == "." {
			for _, val := range cur.children {
				if this.SearchFrom(val, word[i+1:]) {
					return true
				}
			}
			return false
		}

		if _, ok := cur.children[string(c)]; !ok {
			return false
		}
		cur = cur.children[string(c)]
	}
	if cur.isWord {
		return true
	}
	return false
}

func main() {

	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}
