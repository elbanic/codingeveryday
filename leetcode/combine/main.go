
/*
https://leetcode.com/problems/combinations/
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
You may return the answer in any order.

Example 1:
	Input: n = 4, k = 2
	Output:
	[
	  [2,4],
	  [3,4],
	  [2,3],
	  [1,2],
	  [1,3],
	  [1,4],
	]

Example 2:
	Input: n = 1, k = 1
	Output: [[1]]


Constraints:
	1 <= n <= 20
	1 <= k <= n

*/
package main

import "fmt"

func combine(n int, k int) [][]int {

	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	var prev []int
	var comb [][]int
	return combination(k, nums, prev, comb)
}

func combination(depth int, slice []int, prev []int, comb [][]int) [][]int {
	if len(slice) == 0 || depth <= 0 {
		return comb
	}
	depth--
	for i := range slice {
		temp := make([]int, len(prev)+1)
		copy(temp, prev)
		temp[len(temp)-1] = slice[i]
		if depth == 0 {
			comb = append(comb, temp)
		}
		comb = combination(depth, slice[i+1:], temp, comb)
	}
	return comb
}

func main() {
	cb := combine(4, 2)
	for _, v := range cb {
		fmt.Println(v)
	}
}
