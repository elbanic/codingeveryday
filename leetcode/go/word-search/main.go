/*

 */
package main

import (
	"fmt"
)

//dfs

func exist2(board [][]byte, word string) bool {

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if backtrack(board, row, col, word, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, row, col int, word string, index int) bool {
	if index >= len(word) {
		return true
	}

	if row < 0 || row == len(board) || col < 0 || col == len(board[0]) || board[row][col] != word[index] {
		return false
	}

	var ret bool

	board[row][col] = '#'
	rowOffsets := []int{0, 1, 0, -1}
	colOffsets := []int{1, 0, -1, 0}
	for d := 0; d < 4; d++ {
		ret = backtrack(board, row+rowOffsets[d], col+colOffsets[d], word, index+1)
		if ret {
			break
		}
	}
	board[row][col] = word[index]
	return ret
}

type idx struct {
	i int
	j int
}

func exist(board [][]string, word string) bool {

	if len(word) == 0 || len(board) == 0 {
		return false
	}

	var first []idx
	for i, v := range board {
		for j, v2 := range v {
			if string(word[0]) == v2 {
				first = append(first, idx{i, j})
			}
		}
	}

	for _, v := range first {
		visited := make(map[idx]bool)
		visited[v] = true
		if dfs(board, v, word, 0, visited) {
			return true
		}
	}
	return false
}

func dfs(board [][]string, boardId idx, word string, wordId int, visited map[idx]bool) bool {
	wordId++
	visited[boardId] = true
	if wordId > len(word)-1 {
		return true
	}
	neighbors := getNeighbors(boardId.i, boardId.j, board, visited)
	for _, neighbor := range neighbors {
		if board[neighbor.i][neighbor.j] == string(word[wordId]) {
			tmpVisited := make(map[idx]bool)
			for k, v := range visited {
				tmpVisited[k] = v
			}
			if dfs(board, neighbor, word, wordId, visited) {
				return true
			}
			visited = tmpVisited
		}
	}
	return false
}

func getNeighbors(i, j int, board [][]string, visited map[idx]bool) []idx {
	var ret []idx
	neighbor := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, v := range neighbor {
		if !visited[idx{i + v[0], j + v[1]}] && i+v[0] >= 0 && i+v[0] < len(board) && j+v[1] >= 0 && j+v[1] < len(board[0]) {
			ret = append(ret, idx{i + v[0], j + v[1]})
		}
	}
	return ret
}

func main() {
	board := [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	word := "ABCCED"
	fmt.Println(exist(board, word))

	board2 := [][]string{{"A", "B", "C", "E"}, {"S", "F", "E", "S"}, {"A", "D", "E", "E"}}
	word2 := "ABCESEEEFS"
	fmt.Println(exist(board2, word2))

	board3 := [][]string{{"C", "A", "A"}, {"A", "A", "A"}, {"B", "C", "D"}}
	word3 := "AAB"
	fmt.Println(exist(board3, word3))

	board4 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	word4 := "ABCESEEEFS"
	fmt.Println(exist2(board4, word4))
}
