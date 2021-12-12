package main

import "fmt"

func validTree(n int, edges [][]int) bool {

	parent := make(map[int]int)
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	for _, v := range edges {
		if _, ok := parent[v[1]]; ok {
			return false
		}

		parent[v[1]] = v[0]
		union(root, v[0], v[1])
	}

	cur := root[0]
	for _,v := range root{
		if cur != v {
			return false
		}
	}
	return true
}

func union(graph []int, x, y int) {
	rootX := graph[x]
	rootY := graph[y]
	for i,v := range graph {
		if v == rootY {
			graph[i] = rootX
		}
	}
}

func main() {

	n, edges := 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	fmt.Println(validTree(n, edges))

	n2, edges2 := 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	fmt.Println(validTree(n2, edges2))

	n3, edges3 := 2, [][]int{{1, 0}, {2, 0}}
	fmt.Println(validTree(n3, edges3))
}
