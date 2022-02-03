package main

import (
	"fmt"
	"sort"
)

type DisjointSet struct {
	parent []int
	rank   []int
}

func createDisjointSet(n int) DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return DisjointSet{parent, rank}
}

func (ds *DisjointSet) find(i int) int {
	if ds.parent[i] == i {
		return ds.parent[i]
	}
	return ds.find(ds.parent[i])
}

func (ds *DisjointSet) union(i, j int) bool {
	parentI := ds.find(i)
	parentJ := ds.find(j)
	var isMerged bool
	if parentI == parentJ {
		return isMerged
	}

	isMerged = true
	if ds.rank[parentI] > ds.rank[parentJ] {
		ds.parent[parentJ] = parentI
	} else if ds.rank[parentI] < ds.rank[parentJ] {
		ds.parent[parentI] = parentJ
	} else {
		ds.parent[parentI] = parentJ
		ds.rank[parentJ] += 1
	}
	return isMerged
}

func earliestAcq(logs [][]int, n int) int {

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	groupCount := n
	DS := createDisjointSet(n)
	for _, log := range logs {

		if DS.union(log[1], log[2]) {
			groupCount -= 1
		}
		if groupCount == 1 {
			return log[0]
		}
	}
	return -1
}

func main() {
	logs := [][]int{{0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1}}
	n := 4
	fmt.Println(earliestAcq(logs, n))
}
