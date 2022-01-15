/*
https://leetcode.com/problems/jump-game/
You are given an integer array nums. You are initially positioned at the array's first index,
and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.

Example 1:
	Input: nums = [2,3,1,1,4]
	Output: true
	Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
	Input: nums = [3,2,1,0,4]
	Output: false
	Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

package main

import "fmt"

func canJump(nums []int) bool {

	cur := len(nums) - 1
	for cur > 0 {
		max := 0
		for i := cur; i >= 0; i-- {
			if i != len(nums)-1 {
				if len(nums)-1-i <= nums[i] {
					if nums[i] > max {
						cur = i
						max = nums[i]
					}
				}
			}
		}
		if max == 0 {
			return false
		} else {
			nums = nums[:cur+1]
		}
	}
	return true
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))

	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums2))
}
