/*

 */

package main

import (
	"fmt"
	"strconv"
)

const n = 3
const N = 9

var (
	rows         [][]int
	columns      [][]int
	boxes        [][]int
	sudokuSolved bool
	boards       [][]byte
)

func solveSudoku(board [][]byte) {

	boards = board
	sudokuSolved = false
	rows = make([][]int, N)
	for i := range rows {
		rows[i] = make([]int, N+1)
	}
	columns = make([][]int, N)
	for i := range columns {
		columns[i] = make([]int, N+1)
	}
	boxes = make([][]int, N)
	for i := range boxes {
		boxes[i] = make([]int, N+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			num := boards[i][j]
			if num != '.' {
				d, _ := strconv.Atoi(string(num))
				placeNumber(d, i, j)
			}
		}
	}
	backtrackSudoku(0, 0)
}

func couldPlace(d, row, col int) bool {
	idx := (row/n)*n + col/n
	return rows[row][d]+columns[col][d]+boxes[idx][d] == 0
}

func placeNumber(d, row, col int) {
	idx := (row/n)*n + col/n
	rows[row][d]++
	columns[col][d]++
	boxes[idx][d]++
	boards[row][col] = byte(d + '0')
}

func removeNumber(d, row, col int) {
	idx := (row/n)*n + col/n
	rows[row][d]--
	columns[col][d]--
	boxes[idx][d]--
	boards[row][col] = '.'
}

func placeNextNumbers(row, col int) {
	if col == N-1 && row == N-1 {
		sudokuSolved = true
	} else {
		if col == N-1 {
			backtrackSudoku(row+1, 0)
		} else {
			backtrackSudoku(row, col+1)
		}
	}
}

func backtrackSudoku(row, col int) {
	if boards[row][col] == '.' {
		for d := 1; d < 10; d++ {
			if couldPlace(d, row, col) {
				placeNumber(d, row, col)
				placeNextNumbers(row, col)
				if !sudokuSolved {
					removeNumber(d, row, col)
				}
			}
		}
	} else {
		placeNextNumbers(row, col)
	}
}

func printBoard() {
	for i := range boards {
		fmt.Println(boards[i])
	}
}

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	printBoard()
}
