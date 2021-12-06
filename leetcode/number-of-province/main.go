/*
https://leetcode.com/problems/number-of-provinces/
There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b,
and city b is connected directly with city c, then city a is connected indirectly with city c.
A province is a group of directly or indirectly connected cities and no other cities outside of the group.
You are given an n x n matrix isConnected where isConnected[i][j] = 1
if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
Return the total number of provinces.

Example 1:
	Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
	Output: 2

Example 2:
	Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
	Output: 3
*/

package main

import "fmt"

//UnionFind
func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}

	graph := make([]int, len(isConnected))
	for i := range graph {
		graph[i] = i
	}

	for i:=0; i<len(isConnected); i++ {
		for j:=0; j<len(isConnected[0]); j++ {
			if isConnected[i][j] == 1 && i != j{
				graphUnion(graph, i, j)
			}
		}
	}

	ret := make(map[int]bool)
	for _,v := range graph {
		ret[v] = true
	}
	return len(ret)
}

func graphUnion(graph []int, val1, val2 int) {
	root1 := graph[val1]
	root2 := graph[val2]
	for i,v := range graph {
		if v == root2 {
			graph[i] = root1
		}
	}
}

func main() {
	isConnected := [][]int{
		{1,0,0,0,0,0,0,0,0,1,0,0,0,0,0},
		{0,1,0,1,0,0,0,0,0,0,0,0,0,1,0},
		{0,0,1,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,1,0,1,0,0,0,1,0,0,0,1,0,0,0},
		{0,0,0,0,1,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,1,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,1,0,0,0,0,0,0,0,0},
		{0,0,0,1,0,0,0,1,1,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0,0,0},
		{1,0,0,0,0,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0,0,0,0,0,1,0,0,0},
		{0,0,0,0,1,0,0,0,0,0,0,0,1,0,0},
		{0,1,0,0,0,0,0,0,0,0,0,0,0,1,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}}
	fmt.Println(findCircleNum(isConnected))
}
