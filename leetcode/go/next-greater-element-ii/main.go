package main

import "fmt"

func nextGreaterElements(nums []int) []int {

	var ret []int
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if j == len(nums) {
			j = 0
		}
		for j != i {
			if nums[j] > nums[i] {
				break
			}
			j++
			if j == len(nums) {
				j = 0
			}
		}
		if j == i {
			ret = append(ret, -1)
		} else {
			ret = append(ret, nums[j])
		}
	}

	return ret
}

func main() {
	nums := []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}
	fmt.Println(nextGreaterElements(nums))
}
