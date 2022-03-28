package main

import (
	"fmt"
)

func quickSort(nums []int) {
	qSort(nums, 0, len(nums)-1)
}

func qSort(nums []int, lo, hi int) {
	if lo < hi {
		p := partition(nums, lo, hi)
		qSort(nums, lo, p-1)
		qSort(nums, p+1, hi)
	}
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
	}
	tmp := nums[i]
	nums[i] = nums[hi]
	nums[hi] = tmp
	return i
}

func main() {
	nums := []int{1, 5, 3, 2, 8, 7, 6, 4}
	quickSort(nums)
	fmt.Println(nums)
}
