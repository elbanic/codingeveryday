/*
https://leetcode.com/problems/walls-and-gates/
You are given an m x n grid rooms initialized with these three possible values.
	* -1 A wall or an obstacle.
	* 0 A gate.
	* INF Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF
		as you may assume that the distance to a gate is less than 2147483647.
Fill each empty room with the distance to its nearest gate.
If it is impossible to reach a gate, it should be filled with INF.

Example 1:
	Input: rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
	Output: [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]

Example 2:
	Input: rooms = [[-1]]
	Output: [[-1]]

Example 3:
	Input: rooms = [[2147483647]]
	Output: [[2147483647]]

Example 4:
	Input: rooms = [[0]]
	Output: [[0]]
*/

package main

import (
	"fmt"
)

type loc struct {
	i int
	j int
}

func wallsAndGates(rooms [][]int) {
	direction := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				first := loc{i, j}
				queue := []loc{first}
				visited := make(map[loc]bool)

				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					depth := rooms[cur.i][cur.j]

					for _, v := range direction {
						if getNeighbour(cur.i+v[0], cur.j+v[1], rooms, visited) && rooms[cur.i+v[0]][cur.j+v[1]] > 0 && rooms[cur.i+v[0]][cur.j+v[1]] > depth+1 {
							rooms[cur.i+v[0]][cur.j+v[1]] = depth + 1
							queue = append(queue, loc{cur.i + v[0], cur.j + v[1]})
						}
					}
					visited[cur] = true
				}
			}
		}
	}
}
func getNeighbour(i, j int, rooms [][]int, visited map[loc]bool) bool {
	return i >= 0 && j >= 0 && i < len(rooms) && j < len(rooms[0]) && !visited[loc{i, j}]
}

func main() {
	room := [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}
	wallsAndGates(room)

	fmt.Println(room)
}
