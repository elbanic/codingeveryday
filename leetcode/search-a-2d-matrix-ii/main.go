/*
https://leetcode.com/problems/search-a-2d-matrix-ii/
Write an efficient algorithm that searches for a target value in an m x n integer matrix.
The matrix has the following properties:
	Integers in each row are sorted in ascending from left to right.
	Integers in each column are sorted in ascending from top to bottom.

Example 1:
	Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
	Output: true

Example 2:
	Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
	Output: false
*/
package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return searchRec(matrix, target, 0, 0, len(matrix[0])-1, len(matrix)-1)
}

func searchRec(matrix [][]int, target int, left, up, right, down int) bool {
	if left > right || up > down {
		return false
	} else if target < matrix[up][left] || target > matrix[down][right] {
		return false
	}
	mid := left + (right-left)/2
	row := up
	for row <= down && matrix[row][mid] <= target {
		if matrix[row][mid] == target {
			return true
		}
		row++
	}
	return searchRec(matrix, target, left, row, mid-1, down) || searchRec(matrix, target, mid+1, up, right, row-1)
}

func divideConquerMat(matrix [][]int, l, r, t, b int, target int) bool {
	if l == r && t == b {
		return matrix[l][t] == target
	}
	if l+1 >= r && t+1 >= b {
		square := [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
		for _, v := range square {
			if l+v[0] >= 0 && l+v[0] < len(matrix) && t+v[1] >= 0 && t+v[1] < len(matrix[0]) && matrix[l+v[0]][t+v[1]] == target {
				return true
			}
		}
		return false
	}

	midC := l + (r-l)/2
	midR := t + (b-t)/2
	if matrix[midC][midR] == target {
		return true
	} else if matrix[midC][midR] > target {
		// remove right bottom
		if divideConquerMat(matrix, l, midC, t, midR, target) ||
			divideConquerMat(matrix, midC, r, t, midR, target) ||
			divideConquerMat(matrix, l, midC, midR, b, target) {
			return true
		}
	} else if matrix[midC][midR] < target {
		// remove left top
		if divideConquerMat(matrix, midC, r, midR, b, target) ||
			divideConquerMat(matrix, midC, r, t, midR, target) ||
			divideConquerMat(matrix, l, midC, midR, b, target) {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))

	matrix2 := [][]int{{1, 1}}
	target2 := 0
	fmt.Println(searchMatrix(matrix2, target2))
}
