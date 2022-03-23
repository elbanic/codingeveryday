/*

https://leetcode.com/problems/triangle/
Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below.
More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:
	Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
	Output: 11
		Explanation: The triangle looks like:
		   2
		  3 4
		 6 5 7
		4 1 8 3
		The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Example 2:
	Input: triangle = [[-10]]
	Output: -10

Constraints:
	1 <= triangle.length <= 200
	triangle[0].length == 1
	triangle[i].length == triangle[i - 1].length + 1
	-104 <= triangle[i][j] <= 104

*/


package main

import (
	"fmt"
	"strconv"
)

// In-place Algorithm
//func minimumTotal(triangle [][]int) int {
//	for row := 1; row < len(triangle); row++ {
//		for col := 0; col <= row; col++ {
//			smallestAbove := int(^uint(0) >> 1)
//			if col > 0 {
//				smallestAbove = triangle[row-1][col-1]
//			}
//			if col < row {
//				smallestAbove = min(smallestAbove, triangle[row-1][col])
//			}
//			path := smallestAbove + triangle[row][col]
//			triangle[row][col] = path
//		}
//	}
//	return minSlice(triangle[len(triangle)-1])
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func minSlice(s []int) int {
//	if len(s) == 0 {
//		return 0
//	}
//	min := s[0]
//	for i, e := range s {
//		if i == 0 || e < min {
//			min = e
//		}
//	}
//	return min
//}

//Memoization Algorithm
var (
	memoTable map[string]int
)

func minimumTotal(triangle [][]int) int {

	memoTable = make(map[string]int)
	return minPath(triangle, 0,0)
}

func minPath(triangle [][]int, row int, col int) int {
	params := strconv.Itoa(row) + ":" + strconv.Itoa(col)
	if v, ok := memoTable[params]; ok {
		return v
	}
	path := triangle[row][col]
	if row < len(triangle) - 1 {
		path += min(minPath(triangle, row+1, col), minPath(triangle, row+1, col+1))
	}
	memoTable[params] = path
	return path
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	triangle1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle1))
}






