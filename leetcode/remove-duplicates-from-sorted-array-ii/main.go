package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	var k int
	pprev := nums[0]
	prev := nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i] == prev && nums[i] == pprev {
			removeAppend(nums, i)
			i--
			k++
		}
		pprev = prev
		prev = nums[i]
		if i == len(nums)-1-k {
			break
		}
	}
	return len(nums) - k
}

func removeAppend(nums []int, i int) {
	temp := nums[i]
	nums = append(nums[:i], nums[i+1:]...)
	nums = append(nums, temp)
}

func main() {
	nums := []int{1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(nums2)
}
