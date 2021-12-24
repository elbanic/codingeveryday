package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	var mid int
	var found bool
	for left+1 < right && !found {
		mid = left + (right-left)/2
		if nums[mid] == target {
			found = true
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	start, end := -1, -1
	if found {
		start = mid
		end = mid
		for start-1 >= 0 {
			if nums[start-1] == target {
				start--
			} else {
				break
			}
		}
		for end+1 < len(nums) {
			if nums[end+1] == target {
				end++
			} else {
				break
			}
		}
	} else {
		if nums[left] == target {
			start = left
			end = left
		}
		if nums[right] == target {
			if start == -1 {
				start = right
				end = right
			} else {
				end = right
			}
		}
	}
	return []int{start, end}
}

func main() {

	nums := []int{1, 1, 2}
	target := 1
	fmt.Println(searchRange(nums, target))
}
