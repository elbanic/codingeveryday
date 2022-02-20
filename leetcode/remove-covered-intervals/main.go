package main

import (
	"fmt"
	"math"
)

func removeCoveredIntervals(intervals [][]int) int {

	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		intervals = append(intervals[:i], intervals[i+1:]...)
		for j := 0; j < len(intervals); j++ {
			if intervals[j][0] >= cur[0] && intervals[j][1] <= cur[1] {
				intervals = append(intervals[:j], intervals[j+1:]...)
				j = int(math.Max(float64(j-1), -1))
				i = int(math.Max(float64(i-1), -1))
			}
		}
		intervals = append([][]int{cur}, intervals...)
	}
	return len(intervals)
}

func main() {
	intervals := [][]int{{1, 2}, {1, 4}, {3, 4}}
	fmt.Println(removeCoveredIntervals(intervals))
}
