package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {

	var cnt int
	for i := 0; i < len(nums)-2; i++ {
		output := helper(nums, [][]int{}, i, i+1, 0)
		cnt += len(output)
	}

	return cnt
}

func helper(nums []int, output [][]int, start int, end int, diff int) [][]int {
	if end >= len(nums) {
		return output
	}

	if start+1 < end {
		if nums[end]-nums[end-1] == diff {
			output = append(output, nums[start:end+1])
		} else {
			return output
		}
	}

	if start+1 == end {
		diff = nums[end] - nums[start]
	}
	output = helper(nums, output, start, end+1, diff)
	return output
}

func main() {
	nums := []int{1}
	fmt.Println(numberOfArithmeticSlices(nums))
}
