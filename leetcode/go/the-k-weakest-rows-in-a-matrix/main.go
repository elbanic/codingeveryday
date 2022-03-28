package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {

	type rowWeight struct {
		row    int
		weight int
	}
	var rows []rowWeight
	for i, arr := range mat {
		var soldiers int
		for _, val := range arr {
			if val == 1 {
				soldiers++
			}
		}
		rows = append(rows, rowWeight{i, soldiers})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].weight == rows[j].weight {
			return rows[i].row < rows[j].row
		} else {
			return rows[i].weight < rows[j].weight
		}
	})

	var ret []int
	for i := 0; i < k; i++ {
		ret = append(ret, rows[i].row)
	}
	return ret
}

func main() {
	mat := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	k := 3
	fmt.Println(kWeakestRows(mat, k))
}
