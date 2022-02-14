package main

import (
	"fmt"
)

//backtracking
func subsets2(nums []int) [][]int {

	var output [][]int
	output = append(output, []int{})

	for k := 1; k < len(nums)+1; k++ {
		res := backtrack(k, nums, []int{}, [][]int{})
		output = append(output, res...)
	}
	return output
}

func backtrack(depth int, slice []int, prev []int, comb [][]int) [][]int {
	if depth <= 0 {
		comb = append(comb, prev)
		return comb
	}
	depth--
	for i := range slice {
		temp := make([]int, len(prev)+1)
		copy(temp, prev)
		temp[len(temp)-1] = slice[i]
		comb = backtrack(depth, slice[i+1:], temp, comb)
	}
	return comb
}

//cascading
func subsets(nums []int) [][]int {
	temp := helper(0, nums, [][]int{{}})
	return temp
}

func helper(index int, nums []int, output [][]int) [][]int {
	if index >= len(nums) {
		return output
	}

	var newOutput [][]int
	for i := 0; i < len(output); i++ {
		temp := make([]int, len(output[i]))
		copy(temp, output[i])
		temp = append(temp, nums[index])
		newOutput = append(newOutput, temp)
	}

	output = append(output, newOutput...)
	return helper(index+1, nums, output)
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets2(nums))
}
