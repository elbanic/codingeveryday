package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	var ret []int
	sort.Ints(nums)
	for i,v := range nums {
		if v == target {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {

	nums1 := []int{1,2,5,2,3}
	fmt.Println(targetIndices(nums1, 2))

	fmt.Println(targetIndices(nums1, 3))

	fmt.Println(targetIndices(nums1, 5))
	fmt.Println(targetIndices(nums1, 4))
}
