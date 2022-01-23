package main

import (
	"fmt"
	"math"
)

func stoneGameMathematical(piles []int) bool {
	return true
}

func stoneGame(piles []int) bool {
	ret := helper(true, false, piles, []int{}, []int{})
	return ret
}

func helper(alice, bob bool, piles []int, outputA, outputB []int) bool {
	if len(piles) == 2 {
		if alice {
			outputA = append(outputA, int(math.Max(float64(piles[0]), float64(piles[1]))))
		} else {
			outputB = append(outputA, int(math.Max(float64(piles[0]), float64(piles[1]))))
		}
		var sumA, sumB int
		for _, v := range outputA {
			sumA += v
		}
		for _, v := range outputB {
			sumB += v
		}
		return sumA > sumB
	}

	if piles[0] > piles[len(piles)-1] {
		outputA = append(outputA, piles[0])
		return helper(!alice, !bob, piles[1:], outputA, outputB)
	} else {
		outputA = append(outputA, piles[len(piles)-1])
		return helper(!alice, !bob, piles[:len(piles)-1], outputA, outputB)
	}
}

func main() {
	piles := []int{5, 3, 4, 5}
	fmt.Println(stoneGame(piles))
}
