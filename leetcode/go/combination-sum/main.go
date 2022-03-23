package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {

	output := backtrack(candidates, 0, target, 0, []int{}, [][]int{})
	return output
}

func backtrack(candidates []int, index, target, sum int, prev []int, output [][]int) [][]int {
	if sum > target {
		return output
	}

	if sum == target {
		temp := make([]int, len(prev))
		copy(temp, prev)
		output = append(output, temp)
		return output
	}

	for i := index; i < len(candidates); i++ {
		prev = append(prev, candidates[i])
		output = backtrack(candidates, i, target, sum+candidates[i], prev, output)
		prev = prev[:len(prev)-1]
	}
	return output
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
