package main

import (
	"fmt"
	"math"
)

//Using Extra Array
func findMaxLength2(nums []int) int {
	arr := make([]int, 2*len(nums)+1)
	for i := range arr {
		arr[i] = -2
	}
	arr[len(nums)] = -1
	var max, count int

	for i := 0; i < len(nums); i++ {
		var inc int
		if nums[i] == 0 {
			inc = -1
		} else {
			inc = 1
		}
		count += inc
		if arr[count+len(nums)] >= -1 {
			max = int(math.Max(float64(max), float64(i-arr[count+len(nums)])))
		} else {
			arr[count+len(nums)] = i
		}
	}
	return max
}

//brute force
func findMaxLength(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	var max int
	for start := 0; start < len(nums); start++ {
		m := make(map[int]int)

		for end := start; end < len(nums); end++ {
			m[nums[end]]++
			if m[0] == m[1] {
				max = int(math.Max(float64(max), float64(end-start+1)))
			}
		}
	}
	return max
}

func main() {
	nums := []int{0, 1, 0, 1, 1}
	fmt.Println(findMaxLength2(nums))

	nums2 := []int{0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1}
	fmt.Println(findMaxLength2(nums2))
}
