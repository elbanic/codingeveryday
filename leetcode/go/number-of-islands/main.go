/*
https://leetcode.com/problems/number-of-islands/
Given an m x n 2D binary grid which represents a map of '1's (land) and '0's (water),
return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:
	Input: grid = [
	  ["1","1","1","1","0"],
	  ["1","1","0","1","0"],
	  ["1","1","0","0","0"],
	  ["0","0","0","0","0"]
	]
	Output: 1

Example 2:
	Input: grid = [
	  ["1","1","0","0","0"],
	  ["1","1","0","0","0"],
	  ["0","0","1","0","0"],
	  ["0","0","0","1","1"]
	]
	Output: 3
*/

package main

import (
	"fmt"
)

//DFS
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	ret := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				ret++
				dfs(grid, r, c)
			}
		}
	}
	return ret
}

func dfs(grid [][]byte, r int, c int) {
	nr := len(grid)
	nc := len(grid[0])

	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

//BFS
func numIslandsBFS(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	var ret int

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				ret++
				grid[r][c] = '0'
				var neighbors []int
				neighbors = append(neighbors, r*nc+c)
				for len(neighbors) > 0 {
					id := neighbors[len(neighbors)-1]
					neighbors = neighbors[:len(neighbors)-1]
					row := id / nc
					col := id % nc
					if row-1 >= 0 && grid[row-1][col] == '1' {
						neighbors = append(neighbors, (row-1)*nc+col)
						grid[row-1][col] = '0'
					}
					if row+1 < nr && grid[row+1][col] == '1' {
						neighbors = append(neighbors, (row+1)*nc+col)
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						neighbors = append(neighbors, row*nc+col-1)
						grid[row][col-1] = '0'
					}
					if col+1 < nc && grid[row][col+1] == '1' {
						neighbors = append(neighbors, row*nc+col+1)
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}
	return ret
}

func main() {

	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	grid3 := [][]byte{{'1', '0', '1', '1', '0', '1', '1'}}

	fmt.Println(numIslands(grid))
	fmt.Println(numIslands(grid2))
	fmt.Println(numIslands(grid3))
}
