package main

import "fmt"

//divide and conquer
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return search(matrix, target, 0, 0, len(matrix[0])-1, len(matrix)-1)
}

func search(matrix [][]int, target int, left, up, right, down int) bool {
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
	return search(matrix, target, left, row, mid-1, down) || search(matrix, target, mid+1, up, right, row-1)
}

//binary search
func searchMatrix2(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	left, right := 0, len(matrix)*len(matrix[0])-1

	for left <= right {
		mid := left + (right-left)/2
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3}}
	target := 3
	fmt.Println(searchMatrix2(matrix, target))
}
