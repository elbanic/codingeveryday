/*
https://leetcode.com/problems/maximum-distance-in-arrays/
You are given m arrays, where each array is sorted in ascending order.
You can pick up two integers from two different arrays (each array picks one) and calculate the distance.
We define the distance between two integers a and b to be their absolute difference |a - b|.
Return the maximum distance.

Example 1:
	Input: arrays = [[1,2,3],[4,5],[1,2,3]]
	Output: 4
	Explanation: One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.

Example 2:
	Input: arrays = [[1],[1]]
	Output: 0

Example 3:
	Input: arrays = [[1],[2]]
	Output: 1

Example 4:
	Input: arrays = [[1,4],[0,5]]
	Output: 4
*/
package main

import (
	"fmt"
	"math"
)

//brute force
func maxDistance(arrays [][]int) int {
	var mins, maxs []int
	for i := 0; i < len(arrays); i++ {
		mins = append(mins, arrays[i][0])
		maxs = append(maxs, arrays[i][len(arrays[i])-1])
	}
	var maxDist int
	for i := 0; i < len(mins); i++ {
		for j := 0; j < len(maxs); j++ {
			if i != j {
				if maxs[j]-mins[i] > maxDist {
					maxDist = maxs[j] - mins[i]
				}
			}
		}
	}
	return maxDist
}

//single scan
func maxDistance2(arrays [][]int) int {
	var res int
	prevMin := arrays[0][0]
	prevMax := arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		temp1 := math.Abs(float64(arrays[i][len(arrays[i])-1] - prevMin))
		temp2 := math.Abs(float64(prevMax - arrays[i][0]))
		res = int(math.Max(float64(res), math.Max(temp1, temp2)))
		prevMin = int(math.Min(float64(prevMin), float64(arrays[i][0])))
		prevMax = int(math.Max(float64(prevMax), float64(arrays[i][len(arrays[i])-1])))
	}
	return res
}

func main() {
	arr := [][]int{{3, 4, 9}, {1, 10}, {1, 20}}
	fmt.Println(maxDistance2(arr))
}
