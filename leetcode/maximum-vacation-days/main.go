/*
https://leetcode.com/problems/maximum-vacation-days/
*/
package main

import (
	"fmt"
	"math"
)

var MIN_VALUE = 0 << 1

func maxVacationDays(flights [][]int, days [][]int) int {
	memo := make([][]int, len(flights))
	for i := range memo {
		memo[i] = make([]int, len(flights))
		for j := range memo[i] {
			memo[i][j] = MIN_VALUE
		}
	}
	return dfs(flights, days, 0, 0, memo)
}

func dfs(flights [][]int, days [][]int, curCity, weekno int, memo [][]int) int {
	if weekno == len(days[0]) {
		return 0
	}
	if memo[curCity][weekno] != MIN_VALUE {
		return memo[curCity][weekno]
	}
	var maxVac int
	for i := 0; i < len(flights); i++ {
		if flights[curCity][i] == 1 || i == curCity {
			vac := days[i][weekno] + dfs(flights, days, i, weekno+1, memo)
			maxVac = int(math.Max(float64(maxVac), float64(vac)))
		}
	}
	memo[curCity][weekno] = maxVac
	return maxVac
}

func main() {

	flights := [][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}
	days := [][]int{{1, 3, 1}, {6, 0, 3}, {3, 3, 3}}
	fmt.Println(maxVacationDays(flights, days))
}
