package main

import (
	"fmt"
	"math"
	"sort"
)

func findFinalValue(nums []int, original int) int {

	max := nums[0]
	for _, v := range nums {
		max = int(math.Max(float64(v), float64(max)))
	}
	sort.Ints(nums)
	for original <= max {
		if binarySearch(nums, original) {
			original *= 2
		} else {
			return original
		}
	}
	return original
}

func binarySearch(nums []int, k int) bool {

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == k {
			return true
		} else if nums[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	fmt.Println(findFinalValue(nums, original))

	nums2 := []int{2, 7, 9}
	original2 := 4
	fmt.Println(findFinalValue(nums2, original2))
}
