package main

import "fmt"

func rotate(nums []int, k int) {

	if len(nums) <= k {
		if k%len(nums) == 0 {
			return
		}
		k = k % len(nums)
	}

	id := len(nums) - k
	temp := make([]int, k)

	copy(temp, nums[id:])
	temp = append(temp, nums[:id]...)
	copy(nums, temp)
}

func main() {
	nums := []int{1, 2, 3}
	k := 5

	rotate(nums, k)
	fmt.Println(nums)
}
