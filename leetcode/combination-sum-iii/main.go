package main

import "fmt"

func combinationSum3(k int, n int) [][]int {

	var output [][]int
	for i := 1; i < 10; i++ {
		output = append(output, helper(k-1, n-i, []int{i}, [][]int{})...)
	}
	return output
}

func helper(k, n int, prev []int, output [][]int) [][]int {
	if k == 0 && n == 0 {
		temp := make([]int, len(prev))
		copy(temp, prev)
		output = append(output, temp)
		return output
	}
	if k < 0 || n < 0 {
		return output
	}

	for i := prev[len(prev)-1] + 1; i < 10; i++ {
		prev = append(prev, i)
		output = helper(k-1, n-i, prev, output)
		prev = prev[:len(prev)-1]
	}
	return output
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
