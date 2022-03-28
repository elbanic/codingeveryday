package main

import "fmt"

//binary search
func search(nums []int, target int) bool {
	sorted := sort(nums)
	left, right := 0, len(sorted)-1

	for left <= right {
		mid := left + (right-left)/2
		if sorted[mid] == target {
			return true
		} else if sorted[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func sort(nums []int) []int {
	var i int
	for i = 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			if nums[i] > nums[i+1] {
				break
			}
		}
	}
	if i == len(nums) {
		return nums
	}
	return append(nums[i+1:], nums[:i+1]...)
}

func main() {
	nums := []int{1}
	target := 1
	fmt.Println(search(nums, target))
}
