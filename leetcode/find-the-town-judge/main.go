/*
https://leetcode.com/problems/find-the-town-judge/
In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.
If the town judge exists, then:
	The town judge trusts nobody.
	Everybody (except for the town judge) trusts the town judge.
	There is exactly one person that satisfies properties 1 and 2.

You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi.
Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

Example 1:
	Input: n = 2, trust = [[1,2]]
	Output: 2

Example 2:
	Input: n = 3, trust = [[1,3],[2,3]]
	Output: 3

Example 3:
	Input: n = 3, trust = [[1,3],[2,3],[3,1]]
	Output: -1
*/
package main

import "fmt"

//two arrays
func findJudge2(n int, trust [][]int) int {

	if len(trust) < n-1 {
		return -1
	}
	indegrees := make([]int, n+1)
	outdegrees := make([]int, n+1)

	for _, v := range trust {
		outdegrees[v[0]]++
		indegrees[v[1]]++
	}

	for i := 1; i <= n; i++ {
		if indegrees[i] == n-1 && outdegrees[i] == 0 {
			return i
		}
	}
	return -1
}

func findJudge(n int, trust [][]int) int {

	parent := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		parent[i] = i
	}

	for _, v := range trust {
		parent[v[0]] = v[1]
	}

	var couldJudge int
	for i := 1; i < n+1; i++ {
		if parent[i] == i {
			couldJudge = i
		}
	}
	for _, v := range trust {
		if v[1] == couldJudge {
			parent[v[0]] = v[1]
		}
	}
	for i := 1; i < n+1; i++ {
		if parent[i] != couldJudge {
			return -1
		}
	}

	return couldJudge
}
g
func main() {
	n := 3
	trust := [][]int{{1, 2}, {2, 3}}
	fmt.Println(findJudge(n, trust))
}
