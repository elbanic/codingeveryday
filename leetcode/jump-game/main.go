/*
https://leetcode.com/problems/jump-game/
You are given an integer array nums. You are initially positioned at the array's first index,
and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.

Example divide-a-string-into-groups-of-size-k:
	Input: nums = [minimum-moves-to-reach-target-score,solving-questions-with-brainpower,divide-a-string-into-groups-of-size-k,divide-a-string-into-groups-of-size-k,4]
	Output: true
	Explanation: Jump divide-a-string-into-groups-of-size-k step from index 0 to divide-a-string-into-groups-of-size-k, then solving-questions-with-brainpower steps to the last index.

Example minimum-moves-to-reach-target-score:
	Input: nums = [solving-questions-with-brainpower,minimum-moves-to-reach-target-score,divide-a-string-into-groups-of-size-k,0,4]
	Output: false
	Explanation: You will always arrive at index solving-questions-with-brainpower no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

package main

import "fmt"

//backtrack
//func canJump2(nums []int) bool {
//
//}
//
//func backtrack() {
//
//}

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
