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

import (
	"fmt"
	"math"
)

type GOODBAD int

const (
	Good GOODBAD = iota
	Bad
	Unknown
)

//dp bottom up
func canJumpBottomUp(nums []int) bool {
	mem := make([]GOODBAD, len(nums))
	for i := 0; i < len(mem); i++ {
		mem[i] = Unknown
	}
	mem[len(mem)-1] = Good

	for i := len(nums) - 2; i >= 0; i-- {
		furthestJump := int(math.Min(float64(i+nums[i]), float64(len(nums)-1)))
		for j := i + 1; j <= furthestJump; j++ {
			if mem[j] == Good {
				mem[i] = Good
				break
			}
		}
	}
	return mem[0] == Good
}

//memo
var memo []GOODBAD

func canJumpMemo(nums []int) bool {
	memo = make([]GOODBAD, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = Unknown
	}
	memo[len(memo)-1] = Good
	return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(position int, nums []int) bool {
	if memo[position] != Unknown {
		return memo[position] == Good
	}

	furthestJump := int(math.Min(float64(position+nums[position]), float64(len(nums)-1)))
	for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
		if canJumpFromPosition(nextPosition, nums) {
			memo[position] = Good
			return true
		}
	}
	memo[position] = Bad
	return false
}

//backtrack
func canJumpBackTrack(nums []int) bool {
	return backtrack(0, nums)
}

func backtrack(position int, nums []int) bool {
	if position == len(nums)-1 {
		return true
	}

	furthestJump := int(math.Min(float64(position+nums[position]), float64(len(nums)-1)))
	for i := position + 1; i <= furthestJump; i++ {
		if backtrack(i, nums) {
			return true
		}
	}
	return false
}

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
	fmt.Println(canJumpMemo(nums))

	nums2 := []int{1, 5, 2, 1, 0, 2, 0}
	fmt.Println(canJump(nums2))
}
