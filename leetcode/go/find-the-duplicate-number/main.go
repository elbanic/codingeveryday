/*
https://leetcode.com/problems/find-the-duplicate-number/
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
There is only one repeated number in nums, return this repeated number.
You must solve the problem without modifying the array nums and uses only constant extra space.

Example 1:
	Input: nums = [1,3,4,2,2]
	Output: 2

Example 2:
	Input: nums = [3,1,3,4,2]
	Output: 3
*/
package main

import "fmt"

func findDuplicate(nums []int) int {

	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return v
		} else {
			m[v] = true
		}
	}
	return 0
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
