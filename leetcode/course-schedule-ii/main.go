/*
https://leetcode.com/problems/course-schedule-ii/
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first
if you want to take course ai.

	For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them.
If it is impossible to finish all courses, return an empty array.

Example 1:
	Input: numCourses = 2, prerequisites = [[1,0]]
	Output: [0,1]
	Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

Example 2:
	Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
	Output: [0,2,1,3]
	Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
	So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:
	Input: numCourses = 1, prerequisites = []
	Output: [0]
*/

package main

import (
	"fmt"
	"sort"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 1 {
		return []int{0}
	}
	if len(prerequisites) == 0 {
		var ret []int
		for i := 0; i < numCourses; i++ {
			ret = append(ret, i)
		}
		return ret
	}

	sort.Slice(prerequisites, func(i, j int) bool {
		return prerequisites[i][1] < prerequisites[j][1]
	})

	var ret []int
	queue := []int{prerequisites[0][1]}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if !visited[top] {
			visited[top] = true
			ret = append(ret, top)

			for _, v := range prerequisites {
				if v[1] == top {
					queue = append(queue, v[0])
				}
			}
		}
	}

	if !dfs(0, prerequisites, make(map[int]bool)) {
		return []int{}
	}
	if len(ret) < numCourses {
		for i := 0; i < numCourses; i++ {
			var contains bool
			for _, v := range ret {
				if v == i {
					contains = true
					break
				}
			}
			if !contains {
				var cpRet []int
				cpRet = append(cpRet, i)
				for _, v := range ret {
					cpRet = append(cpRet, v)
				}
				ret = cpRet
			}
		}
	}

	return ret
}

func dfs(id int, prerequisites [][]int, visited map[int]bool) bool {
	if visited[id] {
		return false
	}

	visited[id] = true
	for _, v := range prerequisites {
		if id == v[1] {
			return dfs(v[0], prerequisites, visited)
		}
	}
	return true
}

func main() {
	numCourses := 3
	prerequisites := [][]int{{1, 0}}
	fmt.Println(findOrder(numCourses, prerequisites))
}
