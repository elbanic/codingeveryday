/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is
that adjacent houses have security systems connected and it will automatically contact the police
if two adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
	Input: nums = [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
	Total amount you can rob = 1 + 3 = 4.

Example 2:
	Input: nums = [2,7,9,3,1]
	Output: 12
	Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
	Total amount you can rob = 2 + 9 + 1 = 12.
*/

package main

import (
	"fmt"
	"math"
)

//dynamic programming
func rob(nums []int) int {
	N := len(nums)

	if N == 0 {
		return 0
	}

	maxRobbedAmount := make([]int, len(nums)+1)
	maxRobbedAmount[N] = 0
	maxRobbedAmount[N-1] = nums[N-1]

	for i := N - 2; i >= 0; i-- {
		maxRobbedAmount[i] = int(math.Max(float64(maxRobbedAmount[i + 1]), float64(maxRobbedAmount[i + 2] + nums[i])))
	}
	return maxRobbedAmount[0]
}

//recursive + memoization
//func rob(nums []int) int {
//	memo := make([]int, len(nums))
//
//	for i:=0; i<len(memo); i++ {
//		memo[i] = -1
//	}
//
//	return robFrom(0, nums, memo)
//}
//
//func robFrom(i int, nums []int, memo []int) int {
//	if i >= len(nums) {
//		return 0
//	}
//
//	if memo[i] > -1 {
//		return memo[i]
//	}
//
//	ans := int(math.Max(float64(robFrom(i+1, nums, memo)), float64(robFrom(i+2, nums, memo) + nums[i])))
//	memo[i] = ans
//	return ans
//}

func main() {
	num := []int{114, 117, 207, 117, 235, 82, 90, 67, 143,
		146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236,
		81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197,
		187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205,
		169, 241, 202, 144, 240}

	//num := []int{1,2,1,1}
	fmt.Println(rob(num))
}
