/*
https://leetcode.com/problems/n-queens/
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate
a queen and an empty space, respectively.

Example 1:
	Input: n = 4
	Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

Example 2:
	Input: n = 1
	Output: [["Q"]]
 */

package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	board := newBoard(n)
	var boards [][][]int
	boards = findPath(0, board, boards)

	var ret [][]string
	for _,v := range boards {
		ret = append(ret, printBoard(v))
	}
	return ret
}

func findPath(depth int, board [][]int, output [][][]int) [][][]int{
	n := len(board)
	if depth == n {
		output = append(output, board)
		return output
	}
	for y:=0; y<n; y++ {
		if board[depth][y] == 1 {
			cpBoard := copyBoard(board)
			cpBoard = marker(depth, y, cpBoard)
			output = findPath(depth + 1, cpBoard, output)
		}
	}
	return output
}

func marker(x, y int, board [][]int) [][]int{
	n := len(board)
	board[x][y] = -1
	for z:=0; z<n; z++ {
		if board[z][y] != -1 {
			board[z][y] = 0
		}
		if board[x][z] != -1 {
			board[x][z] = 0
		}
		if y+z < n && x+z < n && board[x+z][y+z] != -1{
			board[x+z][y+z] = 0
		}
		if x+z < n && y-z >= 0 && board[x+z][y-z] != -1{
			board[x+z][y-z] = 0
		}
	}
	return board
}

func newBoard(n int) [][]int{
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			board[i][j] = 1
		}
	}
	return board
}

func copyBoard(board [][]int) [][]int{
	n := len(board)
	cpBoard := newBoard(n)
	for i := range board {
		cpBoard[i] = make([]int, n)
		for j := range board[i] {
			cpBoard[i][j] = board[i][j]
		}
	}
	return cpBoard
}

func printBoard(board [][]int) []string{
	var ret []string
	for i:=0; i<len(board); i++ {
		var s string
		for j:=0; j<len(board[i]); j++ {
			switch board[i][j]{
			case 0:
				s += "."
			case -1:
				s += "Q"
			}
		}
		ret = append(ret, s)
	}
	return ret
}

func main() {
	result := solveNQueens(10)
	for _,v := range result {
		fmt.Println(v)
	}
}
