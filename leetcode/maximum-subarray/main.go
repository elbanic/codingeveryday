/*
https://leetcode.com/problems/maximum-subarray/solution/
Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.
A subarray is a contiguous part of an array.

Example 1:
	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
	Input: nums = [1]
	Output: 1

Example 3:
	Input: nums = [5,4,-1,7,8]
	Output: 23
 */

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {

	currentSub := nums[0]
	maxSub := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		currentSub = int(math.Max(float64(num), float64(currentSub + num)))
		maxSub = int(math.Max(float64(maxSub), float64(currentSub)))
	}
	return maxSub
}

func main() {
	nums := []int{-1}
	fmt.Println(maxSubArray(nums))
}
