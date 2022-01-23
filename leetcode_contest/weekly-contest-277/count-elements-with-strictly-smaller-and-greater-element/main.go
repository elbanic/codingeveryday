package main

import (
	"fmt"
	"sort"
)

func countElements(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	if min == max {
		return 0
	}

	var ret []int
	for _, v := range nums {
		if min < v && v < max {
			ret = append(ret, v)
		}
	}
	return len(ret)
}

func main() {
	nums := []int{-3, 3, 3, 90}
	fmt.Println(countElements(nums))
}
