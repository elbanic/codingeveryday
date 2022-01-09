package main

import "fmt"

func checkValid(matrix [][]int) bool {

	m := make(map[int]bool)
	var status bool
	for i := 1; i < len(matrix)+1; i++ {
		m[i] = status
	}

	for i := 0; i < len(matrix); i++ {
		status = !status
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] > len(matrix) || matrix[i][j] < 1 {
				return false
			}
			if m[matrix[i][j]] == status {
				return false
			}
			m[matrix[i][j]] = status
		}
	}

	status = false
	for i := 1; i < len(matrix)+1; i++ {
		m[i] = status
	}
	for i := 0; i < len(matrix); i++ {
		status = !status
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[j][i] > len(matrix) || matrix[j][i] < 1 {
				return false
			}
			if m[matrix[j][i]] == status {
				return false
			}
			m[matrix[j][i]] = status
		}
	}
	return true
}

func main() {
	matrix := [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}
	fmt.Println(checkValid(matrix))

	matrix2 := [][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println(checkValid(matrix2))
}
