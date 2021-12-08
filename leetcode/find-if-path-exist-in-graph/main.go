/*
https://leetcode.com/problems/find-if-path-exists-in-graph/
There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive).
The edges in the graph are represented as a 2D integer array edges,
where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi.
Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.

You want to determine if there is a valid path that exists from vertex start to vertex end.
Given edges and the integers n, start, and end, return true if there is a valid path from start to end,
or false otherwise.

Example 1:
	Input: n = 3, edges = [[0,1],[1,2],[2,0]], start = 0, end = 2
	Output: true
	Explanation: There are two paths from vertex 0 to vertex 2:
	- 0 → 1 → 2
	- 0 → 2

Example 2:
	Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], start = 0, end = 5
	Output: false
	Explanation: There is no path from vertex 0 to vertex 5.
*/

package main

import (
	"fmt"
)

// DFS
var stack []int
var visited map[int]bool

func validPath(n int, edges [][]int, start int, end int) bool {

	if len(edges) == 0 {
		if start == end && start == 0 {
			return true
		}
		return false
	}

	for _,v := range edges {
		if (v[0]==start && v[1]==end) || (v[0]==end && v[1]==start){
			return true
		}
	}

	stack = []int{edges[0][0]}
	visited = make(map[int]bool)
	traverseVertices(edges)

	startExist, ok1 := visited[start]
	endExist, ok2 := visited[end]
	return ok1 && ok2 && startExist && endExist
}

func traverseVertices(edges [][]int) {
	if len(stack) == 0 {
		return
	}
	cur := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	if _, ok := visited[cur]; !ok {
		visited[cur] = true
		for _, v := range edges {
			if v[0] == cur {
				stack = append(stack, v[1])
			}
			if v[1] == cur {
				stack = append(stack, v[0])
			}
		}
	}
	traverseVertices(edges)
}

//disjoint set
func validPathDS(n int, edges [][]int, start int, end int) bool {

	graph := make([]int, n)
	for i := range graph {
		graph[i] = i
	}

	for _, v := range edges {
		union(graph, v[0], v[1])
	}
	return find(graph, start) == find(graph, end)
}

func union(graph []int, x, y int){

	rootX := graph[x]
	rootY := graph[y]

	for i,v := range graph {
		if v == rootY {
			graph[i] = rootX
		}
	}
}

func find(graph []int, x int) int{
	return graph[x]
}

func main() {
	n := 10
	edges := [][]int{{4,3},{1,4},{4,8},{1,7},{6,4},{4,2},{7,4},{4,0},{0,9},{5,4}}
	start := 5
	end := 9
	fmt.Println(validPathDS(n, edges, start, end))
}
