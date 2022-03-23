/*
https://leetcode.com/problems/permutations/
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
	Input: nums = [1,2,3]
	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
	Input: nums = [0,1]
	Output: [[0,1],[1,0]]

Example 3:
	Input: nums = [1]
	Output: [[1]]
*/
package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var ret [][]int

	newNums := make([]int, len(nums))
	copy(newNums, nums)

	bactrack(len(nums), newNums, &ret, 0)
	return ret
}

func bactrack(n int, nums []int, output *[][]int, first int) {
	if first == n {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*output = append(*output, temp)
	}

	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		bactrack(n, nums, output, first+1)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(permutations(nums2))
}
