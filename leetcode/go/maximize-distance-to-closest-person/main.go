/*
https://leetcode.com/problems/maximize-distance-to-closest-person/submissions/
You are given an array representing a row of seats where seats[i] = 1 represents a person sitting in the ith seat,
and seats[i] = 0 represents that the ith seat is empty (0-indexed).

There is at least one empty seat, and at least one person sitting.
Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to the closest person.

Example 1:
	Input: seats = [1,0,0,0,1,0,1]
	Output: 2
	Explanation:
	If Alex sits in the second open seat (i.e. seats[2]), then the closest person has distance 2.
	If Alex sits in any other open seat, the closest person has distance 1.
	Thus, the maximum distance to the closest person is 2.

Example 2:
	Input: seats = [1,0,0,0]
	Output: 3
	Explanation:
	If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats away.
	This is the maximum distance possible, so the answer is 3.

Example 3:
	Input: seats = [0,1]
	Output: 1
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func maxDistToClosest(seats []int) int {

	leftFilled := seats[0] == 1
	rightFilled := seats[len(seats)-1] == 1

	n := len(seats)
	var max, firstEnd int
	start, end := -1, -1
	i, j := 0, 0

	for i < n && j < n {
		if seats[i] == 0 {
			start = i
		} else {
			i++
		}

		if seats[j] == 0 {
			end = j
		} else {
			if start == 0 {
				firstEnd = end
			}
			if end-start > max {
				max = end - start
			}
			i = j
		}
		j++
	}

	var retArr []int
	left, right, mid := -1, -1, max/2+1
	if !leftFilled {
		if firstEnd+1 > max/2 {
			left = firstEnd + 1
		}
	}
	if !rightFilled {
		if j-i > max/2 {
			right = j - i
		}
	}
	retArr = append(retArr, left)
	retArr = append(retArr, right)
	retArr = append(retArr, mid)
	sort.Ints(retArr)

	return retArr[len(retArr)-1]
}

//Next Array
func maxDistToClosest2(seats []int) int {
	n := len(seats)
	left, right := make([]int, n), make([]int, n)
	for i := range left {
		left[i] = n
		right[i] = n
	}

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			left[i] = 0
		} else if i > 0 {
			left[i] = left[i-1] + 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		if seats[i] == 1 {
			right[i] = 0
		} else if i < n-1 {
			right[i] = right[i+1] + 1
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		if seats[i] == 0 {
			ans = int(math.Max(float64(ans), math.Min(float64(left[i]), float64(right[i]))))
		}
	}
	return ans
}

//Two Pointer
func maxDistToClosest3(seats []int) int {
	n := len(seats)
	prev, future := -1, 0
	var ans int

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			prev = i
		} else {
			for future < n && seats[future] == 0 || future < i {
				future++
			}

			left := n
			if prev != -1 {
				left = i - prev
			}
			right := n
			if future != n {
				right = future - i
			}
			ans = int(math.Max(float64(ans), math.Min(float64(left), float64(right))))
		}
	}
	return ans
}

func main() {
	seats := []int{1, 0, 0, 0}
	fmt.Println(maxDistToClosest3(seats))

	seats2 := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0}
	fmt.Println(maxDistToClosest3(seats2))
}
