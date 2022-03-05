package main

import (
	"fmt"
)

func deleteAndEarn(nums []int) int {

	out := helper(nums, 0, -1, make(map[int]int))
	return out
}

func helper(nums []int, sum int, key int, visited map[int]int) int {
	if len(nums) == 0 {
		return sum
	}

	var max int
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		var temp []int
		for j := 0; j < len(nums); j++ {
			if i != j && n+1 != nums[j] && n-1 != nums[j] {
				temp = append(temp, nums[j])
			}
		}
		var res int
		if val, ok := visited[key]; ok {
			res = val
		} else {
			res = helper(temp, sum+n, nums[i], visited)
			visited[nums[i]] = res
		}

		if res > max {
			max = res
		}
	}
	return max
}

func main() {
	nums := []int{8, 3, 4, 7, 6, 6, 9, 2, 5, 8, 2, 4, 9, 5, 9, 1, 5, 7, 1, 4}
	fmt.Println(deleteAndEarn(nums))
}
