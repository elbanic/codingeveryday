/*
https://leetcode.com/problems/find-k-closest-elements/
Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array.
The result should also be sorted in ascending order.

An integer a is closer to x than an integer b if:
	|a - x| < |b - x|, or
	|a - x| == |b - x| and a < b

Example 1:
	Input: arr = [1,2,3,4,5], k = 4, x = 3
	Output: [1,2,3,4]

Example 2:
	Input: arr = [1,2,3,4,5], k = 4, x = -1
	Output: [1,2,3,4]
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {

	left := 0
	right := len(arr) - 1
	var mid int
	var found bool
	for left+1 < right && !found {
		mid = left + (right-left)/2
		if arr[mid] == x {
			found = true
		} else if arr[mid] > x {
			right = mid
		} else {
			left = mid
		}
	}

	var ret []int
	if found {
		ret = append(ret, arr[mid])

		left = mid - 1
		right = mid + 1
	}

	for len(ret) < k {
		if left < 0 {
			ret = append(ret, arr[right])
			right++
		} else if right > len(arr)-1 {
			ret = append(ret, arr[left])
			left--
		} else {
			a := math.Abs(float64(x - arr[left]))
			b := math.Abs(float64(x - arr[right]))
			if a > b {
				ret = append(ret, arr[right])
				right++
			} else {
				ret = append(ret, arr[left])
				left--
			}
		}
	}
	sort.Ints(ret)
	return ret
}

func main() {
	arr := []int{1, 1, 1, 10, 10, 10}
	k := 1
	x := 9
	fmt.Println(findClosestElements(arr, k, x))
}
