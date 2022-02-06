package main

import (
	"fmt"
	"sort"
)

func sortEvenOdd(nums []int) []int {

	var even, odd []int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}

	sort.Ints(even)
	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})

	var ret []int
	var ev, od int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ret = append(ret, even[ev])
			ev++
		} else {
			ret = append(ret, odd[od])
			od++
		}
	}
	return ret
}

func main() {

	nums := []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}
	fmt.Println(sortEvenOdd(nums))
}
