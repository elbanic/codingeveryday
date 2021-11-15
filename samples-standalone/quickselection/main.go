package main

import (
	"fmt"
)

func partition(nums []int, left, right, pivotIndex int) int {

	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storeIndex := left

	for i:=left; i<=right; i++ {
		if nums[i] < pivot {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}
	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}

func quickSelect(nums []int, left, right, k_smallest int) int {

	if left == right {
		return nums[left]
	}

	pivotIndex := 0
	pivotIndex = partition(nums, left, right, pivotIndex)

	if k_smallest == pivotIndex {
		return nums[pivotIndex]
	} else if k_smallest < pivotIndex {
		return quickSelect(nums, left, pivotIndex - 1, k_smallest)
	} else {
		return quickSelect(nums, pivotIndex + 1, right, k_smallest)
	}
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums) - 1, len(nums) - k)
}

func main() {
	nums1, k := []int{3,2,1,5,6,4}, 2
	fmt.Println(findKthLargest(nums1, k)) //5

	nums2, k := []int{3,2,3,1,2,4,5,5,6}, 4
	fmt.Println(findKthLargest(nums2, k)) //4
}
