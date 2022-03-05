package main

import (
	"fmt"
)

func mostFrequent(nums []int, key int) int {

	var max, maxFreq int
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
			if m[nums[i+1]] > maxFreq {
				maxFreq = m[nums[i+1]]
				max = nums[i+1]
			}
		}
	}
	return max
}

func main() {
	nums := []int{1, 555, 2, 555, 3, 555, 2}
	key := 555
	fmt.Println(mostFrequent(nums, key))
}
