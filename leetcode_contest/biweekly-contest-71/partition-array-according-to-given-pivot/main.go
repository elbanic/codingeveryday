package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {

	var ret1 []int
	var ret2 []int
	var ret3 []int
	for _, v := range nums {
		if v < pivot {
			ret1 = append(ret1, v)
		} else if v == pivot {
			ret2 = append(ret2, v)
		} else {
			ret3 = append(ret3, v)
		}
	}

	ret1 = append(ret1, ret2...)
	ret1 = append(ret1, ret3...)
	return ret1
}

func main() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10
	fmt.Println(pivotArray(nums, pivot))

	nums2 := []int{-3, 4, 3, 2}
	pivot2 := 2
	fmt.Println(pivotArray(nums2, pivot2))
}
