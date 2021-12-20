/*
https://leetcode.com/problems/minimum-absolute-difference/
Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.
Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
	a, b are from arr
	a < b
	b - a equals to the minimum absolute difference of any two elements in arr

Example 1:
	Input: arr = [4,2,1,3]
	Output: [[1,2],[2,3],[3,4]]
	Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

Example 2:
	Input: arr = [1,3,6,10,15]
	Output: [[1,3]]

Example 3:
	Input: arr = [3,8,-10,23,19,-4,-14,27]
	Output: [[-14,-10],[19,23],[23,27]]

*/
package main

import (
	"fmt"
	"math"
	"sort"
)

type pair struct {
	i int
	j int
}

func minimumAbsDifference(arr []int) [][]int {

	sort.Ints(arr)
	pairs := make(map[pair]int)
	var min int

	for i := 0; i < len(arr)-1; i++ {
		j := i + 1
		abs := int(math.Abs(float64(arr[i] - arr[j])))
		if i == 0 {
			min = abs
		} else {
			if min > abs {
				min = abs
			}
		}

		if min == abs {
			pairs[pair{arr[i], arr[j]}] = abs
		}
	}

	var ret [][]int
	for k, v := range pairs {
		if v == min {
			ret = append(ret, []int{k.i, k.j})
		}
	}
	sort.Slice(ret, func(i int, j int) bool {
		return ret[i][0] < ret[j][0]
	})
	return ret
}

func main() {
	arr := []int{1, 3, 6, 10, 15}
	fmt.Println(minimumAbsDifference(arr))

	arr2 := []int{3, 8, -10, 23, 19, -4, -14, 27}
	fmt.Println(minimumAbsDifference(arr2))
}
