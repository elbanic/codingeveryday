/*
https://leetcode.com/problems/kth-missing-positive-number/
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
Find the kth positive integer that is missing from this array.

Example 1:
	Input: arr = [2,3,4,7,11], k = 5
	Output: 9
	Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

Example 2:
	Input: arr = [1,2,3,4], k = 2
	Output: 6
	Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
*/
package main

import "fmt"

//brute force
func findKthPositive(arr []int, k int) int {

	if len(arr) == 0 {
		return k
	}
	last := arr[len(arr)-1]
	if last-len(arr) >= k {
		var newArr []int
		for i := 1; i < last; i++ {
			if !contains(arr, i) {
				newArr = append(newArr, i)
			}
		}
		return newArr[k-1]
	}
	return k + len(arr)
}

//binary search
func contains(nums []int, e int) bool {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == e {
			return true
		} else if nums[mid] > e {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

func main() {
	arr := []int{2}
	k := 1
	fmt.Println(findKthPositive(arr, k))
}
