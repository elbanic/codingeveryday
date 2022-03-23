/*
https://leetcode.com/problems/cherry-pickup-ii/
You are given a rows x cols matrix grid representing a field of cherries where grid[i][j] represents the number of cherries that you can collect from the (i, j) cell.

You have two robots that can collect cherries for you:
	* Robot #1 is located at the top-left corner (0, 0), and
	* Robot #2 is located at the top-right corner (0, cols - 1).

Return the maximum number of cherries collection using both robots by following the rules below:
	* From a cell (i, j), robots can move to cell (i + 1, j - 1), (i + 1, j), or (i + 1, j + 1).
	* When any robot passes through a cell, It picks up all cherries, and the cell becomes an empty cell.
	* When both robots stay in the same cell, only one takes the cherries.
	* Both robots cannot move outside of the grid at any moment.
	* Both robots should reach the bottom row in grid.
*/
package main

import (
	"fmt"
	"math"
)

func cherryPickup(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cache := make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, n)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				cache[i][j][k] = -1
			}
		}
	}
	return dp(0, 0, n-1, grid, cache)
}

func dp(row, col1, col2 int, grid [][]int, cache [][][]int) int {
	if col1 < 0 || col1 >= len(grid[0]) || col2 < 0 || col2 >= len(grid[0]) {
		return 0
	}
	if cache[row][col1][col2] != -1 {
		return cache[row][col1][col2]
	}
	res := 0
	res += grid[row][col1]
	if col1 != col2 {
		res += grid[row][col2]
	}
	if row != len(grid)-1 {
		max := 0
		for newCol1 := col1 - 1; newCol1 <= col1+1; newCol1++ {
			for newCol2 := col2 - 1; newCol2 <= col2+1; newCol2++ {
				max = int(math.Max(float64(max), float64(dp(row+1, newCol1, newCol2, grid, cache))))
			}
		}
		res += max
	}
	cache[row][col1][col2] = res
	return res
}

func main() {
	grid := [][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}
	fmt.Println(cherryPickup(grid))
}
