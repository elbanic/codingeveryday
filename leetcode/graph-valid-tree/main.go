package main

import (
	"fmt"
)

//union find
type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{make([]int, n)}
	uf.makeset()
	return &uf
}

func (uf *UnionFind) makeset() {
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

func (uf UnionFind) find(A int) int {
	for uf.parent[A] != A {
		A = uf.parent[A]
	}
	return A
}

func (uf *UnionFind) union(A, B int) bool {

	rootA := uf.find(A)
	rootB := uf.find(B)

	if rootA == rootB {
		return false
	}
	uf.parent[rootB] = rootA
	return true
}

func validTree(n int, edges [][]int) bool {

	if len(edges) != n-1 {
		return false
	}

	var uf *UnionFind
	uf = NewUnionFind(n)

	for _, v := range edges {
		A := v[0]
		B := v[1]

		if !uf.union(A, B) {
			return false
		}
	}
	return true
}

func main() {

	n, edges := 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	fmt.Println(validTree(n, edges))

	n2, edges2 := 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}}
	fmt.Println(validTree(n2, edges2))
}
