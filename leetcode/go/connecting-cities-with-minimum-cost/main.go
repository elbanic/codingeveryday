/*
https://leetcode.com/problems/connecting-cities-with-minimum-cost/
There are n cities labeled from 1 to n. You are given the integer n and an array connections
where connections[i] = [xi, yi, costi] indicates that the cost of connecting city xi and city yi (bidirectional connection) is costi.
Return the minimum cost to connect all the n cities such that there is at least one path between each pair of cities.
If it is impossible to connect all the n cities, return -1,
The cost is the sum of the connections' costs used.

Example 1:
	Input: n = 3, connections = [[1,2,5],[1,3,6],[2,3,1]]
	Output: 6
	Explanation: Choosing any 2 edges will connect all cities so we choose the minimum 2.

Example 2:
	Input: n = 4, connections = [[1,2,3],[3,4,4]]
	Output: -1
	Explanation: There is no way to connect all cities even if all edges are used.
*/
package main

import (
	"fmt"
	"sort"
)

//Minimum Spanning Tree (Using Kruskal's algorithm)
func minimumCost2(n int, connections [][]int) int {
	rank := make([]int, n+1)
	parents := make([]int, n+1)
	for i := 0; i <= n; i++ {
		rank[i] = i
		parents[i] = i
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	var total, cost int

	for i := 0; i < len(connections); i++ {
		a := connections[i][0]
		b := connections[i][1]
		findA := find(parents, a)
		findB := find(parents, b)
		if findA == findB {
			continue
		}
		union(rank, parents, a, b)
		cost += connections[i][2]
		total++
	}
	if total == n-1 {
		return cost
	}
	return -1
}

func find(parents []int, a int) int {
	for a != parents[a] {
		parents[a] = parents[parents[a]]
		a = parents[a]
	}
	return a
}

func union(rank []int, parents []int, a int, b int) {
	rootA := find(parents, a)
	rootB := find(parents, b)
	if rootA == rootB {
		return
	}
	if rank[rootA] > rank[rootB] {
		parents[rootB] = rootA
		rank[rootA] += rank[rootB]
	} else {
		parents[rootA] = rootB
		rank[rootB] += rank[rootA]
	}
}

//BFS
type pair struct {
	i int
	j int
	c int
}

func minimumCost(n int, connections [][]int) int {

	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	var minCost int
	first := pair{connections[0][0], connections[0][1], connections[0][2]}
	queue := []pair{first}
	visited := make(map[int]bool)
	visited[first.i] = true
	visited[first.j] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		minCost += cur.c
		for i := 0; i < len(connections); i++ {
			match := pair{connections[i][0], connections[i][1], connections[i][2]}
			if !visited[match.i] || !visited[match.j] {
				if cur.i == match.i || cur.i == match.j || cur.j == match.i || cur.j == match.j {
					visited[match.i] = true
					visited[match.j] = true
					queue = append(queue, match)
				}
			}
		}
	}

	if len(visited) != n {
		return -1
	}
	return minCost
}

func main() {
	n := 3
	connections := [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}
	fmt.Println(minimumCost2(n, connections))

	n2 := 7
	connections2 := [][]int{{2, 1, 87129}, {3, 1, 14707}, {4, 2, 34505}, {5, 1, 71766}, {6, 5, 2615}, {7, 2, 37352}}
	fmt.Println(minimumCost2(n2, connections2))
}
