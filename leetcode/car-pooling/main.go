/*
https://leetcode.com/problems/car-pooling/
There is a car with capacity empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).
You are given the integer capacity and an array trips where trip[i] = [numPassengersi, fromi, toi] indicates
that the ith trip has numPassengersi passengers and the locations to pick them up and drop them off are fromi and toi respectively.
The locations are given as the number of kilometers due east from the car's initial location.
Return true if it is possible to pick up and drop off all passengers for all the given trips, or false otherwise.

Example 1:
	Input: trips = [[2,1,5],[3,3,7]], capacity = 4
	Output: false

Example 2:
	Input: trips = [[2,1,5],[3,3,7]], capacity = 5
	Output: true
*/
package main

import (
	"fmt"
	"sort"
)

//using timestamp
func carPooling2(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return false
	}
	timestamp := make(map[int]int)
	for _, trip := range trips {
		if v, ok := timestamp[trip[1]]; ok {
			timestamp[trip[1]] = v + trip[0]
		} else {
			timestamp[trip[1]] = trip[0]
		}
		if v, ok := timestamp[trip[2]]; ok {
			timestamp[trip[2]] = v - trip[0]
		} else {
			timestamp[trip[2]] = -trip[0]
		}
	}

	var res [][]int
	for key, val := range timestamp {
		res = append(res, []int{key, val})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})

	usedCap := 0
	for _, v := range res {
		usedCap += v[1]
		if usedCap > capacity {
			return false
		}
	}
	return true
}

//brute force
func carPooling(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return false
	}
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][2] < trips[j][2]
	})

	max := trips[len(trips)-1][2]
	var totalPass int
	for i := 0; i < max; i++ {
		for _, v := range trips {
			if v[1] <= i && i < v[2] {
				totalPass += v[0]
			}
		}
		if totalPass > capacity {
			return false
		}
		totalPass = 0
	}
	return true
}

func main() {
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 4
	fmt.Println(carPooling2(trips, capacity))

	trips2 := [][]int{{7, 5, 6}, {6, 7, 8}, {10, 1, 6}}
	capacity2 := 16
	fmt.Println(carPooling2(trips2, capacity2))
}
