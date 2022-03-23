package main

import (
	"fmt"
	"sort"
	"strings"
)

type TreeNode struct {
	isDir   bool
	name    string
	content string
	leaves  []*TreeNode
}

type FileSystem struct {
	root *TreeNode
}

func Constructor() FileSystem {
	root := &TreeNode{true, "", "", []*TreeNode{}}
	return FileSystem{root}
}

func (this *FileSystem) Ls(path string) []string {

	node := this.find(path)
	if node != nil {
		if node.isDir {
			var ret []string
			for _, leaf := range node.leaves {
				ret = append(ret, leaf.name)
			}
			sort.Strings(ret)
			return ret
		} else {
			return []string{node.name}
		}
	}
	return nil
}

func (this *FileSystem) Mkdir(path string) {
	paths := strings.Split(path, "/")
	if paths[0] == "" {
		paths = paths[1:]
	}

	cur := this.root
	for _, v := range paths {
		var found bool
		for j := range cur.leaves {
			if cur.leaves[j].isDir && cur.leaves[j].name == v {
				cur = cur.leaves[j]
				found = true
				break
			}
		}
		if !found {
			next := &TreeNode{true, v, "", []*TreeNode{}}
			cur.leaves = append(cur.leaves, next)
			cur = next
		}
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {

	node := this.find(filePath)
	if node != nil {
		node.content += content
	} else {
		paths := strings.Split(filePath, "/")
		path := strings.Join(paths[:len(paths)-1], "/")
		this.Mkdir(path)
		parent := this.find(path)
		parent.leaves = append(parent.leaves, &TreeNode{false, paths[len(paths)-1], content, nil})
	}
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	node := this.find(filePath)
	if node != nil && !node.isDir {
		return node.content
	}
	return ""
}

func (this *FileSystem) find(path string) *TreeNode {
	paths := strings.Split(path, "/")
	if paths[0] == "" {
		paths = paths[1:]
	}
	if len(paths) == 1 && paths[0] == "" {
		return this.root
	}

	cur := this.root
	for _, v := range paths {
		var found bool
		for j := range cur.leaves {
			if cur.leaves[j].name == v {
				cur = cur.leaves[j]
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return cur
}

func main() {

	//["FileSystem","mkdir","ls","ls","mkdir","ls","ls","addContentToFile","readContentFromFile","ls","readContentFromFile"]
	//[[],["/zijzllb"],["/"],["/zijzllb"],["/r"],["/"],["/r"],["/zijzllb/hfktg","d"],["/zijzllb/hfktg"],["/"],["/zijzllb/hfktg"]]

	fs := Constructor()
	fs.Mkdir("/zijzllb")
	fmt.Println(fs.Ls("/"))
	fmt.Println(fs.Ls("/zijzllb"))
	fs.Mkdir("/r")
	fmt.Println(fs.Ls("/"))
	fmt.Println(fs.Ls("/r"))
	fs.AddContentToFile("/zijzllb/hfktg", "d")
	fs.ReadContentFromFile("/zijzllb/hfktg")
	fmt.Println(fs.Ls("/"))
	fmt.Println(fs.ReadContentFromFile("/zijzllb/hfktg"))

	fs1 := Constructor()
	fmt.Println(fs1.Ls("/"))
	fs1.Mkdir("/a/b/c")
	fs1.AddContentToFile("/a/b/c/d", "hello")
	fmt.Println(fs1.Ls("/a"))
	fmt.Println(fs1.ReadContentFromFile("/a/b/c/d"))

}
