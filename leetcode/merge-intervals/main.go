/*
https://leetcode.com/problems/merge-intervals/
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:
	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
	Input: intervals = [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i := 0
	for i < len(intervals)-1 {
		j := i + 1
		cur := intervals[i]
		next := intervals[j]

		if cur[1] >= next[0] {
			if cur[1] < next[1] {
				cur[1] = next[1]
				temp := append(intervals[:i], cur)
				intervals = append(temp, intervals[j+1:]...)
			} else if cur[1] >= next[1] {
				temp := append(intervals[:i], cur)
				intervals = append(temp, intervals[j+1:]...)
			}
		} else {
			i++
		}
	}
	return intervals
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))

	intervals2 := [][]int{{1, 6}, {2, 3}}
	fmt.Println(merge(intervals2))

	intervals3 := [][]int{{1, 3}, {3, 6}}
	fmt.Println(merge(intervals3))

	intervals4 := [][]int{{1, 3}, {4, 6}}
	fmt.Println(merge(intervals4))
}
