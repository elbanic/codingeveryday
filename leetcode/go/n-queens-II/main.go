package main

import "fmt"

func totalNQueens(n int) int {
	result := solveNQueens(n)
	return len(result)
}

func solveNQueens(n int) [][][]int {
	//board
	//	0 is default
	//	1 is queen
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	return backtrackQueen(board, 0, n, [][][]int{})
}

func backtrackQueen(board [][]int, row, queen int, output [][][]int) [][][]int {
	if queen == 0 {
		fmt.Println(board)
		temp := make([][]int, len(board))
		for i := range temp {
			temp[i] = make([]int, len(board))
		}
		copy(temp, board)
		output = append(output, temp)
		return output
	}
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			placeQueen(board, row, col)
			output = backtrackQueen(board, row+1, queen-1, output)
			removeQueen(board, row, col)
		}
	}
	return output
}

func isValid(board [][]int, i, j int) bool {
	var ret bool
	if i >= 0 && i < len(board) && j >= 0 && j < len(board) {
		ret = board[i][j] == 0
		if !ret {
			return false
		}
		for x := 0; x < len(board); x++ {
			if board[i][x] != 0 || board[x][j] != 0 {
				return false
			}
			if i-x >= 0 && j-x >= 0 && board[i-x][j-x] != 0 {
				return false
			}
			if i-x >= 0 && j+x < len(board) && board[i-x][j+x] != 0 {
				return false
			}
			if i+x < len(board) && j+x < len(board) && board[i+x][j+x] != 0 {
				return false
			}
			if i+x < len(board) && j-x >= 0 && board[i+x][j-x] != 0 {
				return false
			}
		}
	}
	return ret
}

func placeQueen(board [][]int, i, j int) {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
		board[i][j] = 1
	}
}

func removeQueen(board [][]int, i, j int) {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
		board[i][j] = 0
	}
}

func main() {
	fmt.Println(totalNQueens(4))
}
