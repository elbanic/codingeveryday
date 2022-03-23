package main

import "fmt"

//binary search
func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))

	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	fmt.Println(searchInsert(nums2, target2))
}
