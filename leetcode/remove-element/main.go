package main

import "fmt"

func removeElement(nums []int, val int) int {

	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	fmt.Println(removeElement(nums2, val2))
	fmt.Println(nums2)
}
