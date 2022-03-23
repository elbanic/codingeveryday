package main

import (
	"fmt"
	"math"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	glass := make([][]float64, 102)
	for i := range glass {
		glass[i] = make([]float64, i+1)
	}
	glass[0][0] = float64(poured)
	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			q := (glass[r][c] - 1) / 2
			if q > 0 {
				glass[r+1][c] += q
				glass[r+1][c+1] += q
			}
		}
	}
	return math.Min(1., glass[query_row][query_glass])
}

func main() {
	fmt.Println(champagneTower(7, 3, 1))
}
