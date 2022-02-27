package main

import (
	"fmt"
	"sort"
)

//brute force
func minimumTime(time []int, totalTrips int) int64 {
	min := time[0]
	for i := range time {
		if min > time[i] {
			min = time[i]
		}
	}
	MaxInt := int(^uint32(0)>>1) - 1
	for i := min; i < MaxInt; i++ {
		var newTotalTrips int
		for j := range time {
			newTotalTrips += i / time[j]
		}
		if newTotalTrips >= totalTrips {
			return int64(i)
		}
	}
	return 0
}

//binary search
func minimumTime2(time []int, totalTrips int) int64 {
	sort.Ints(time)
	left := time[0]
	right := int(^uint32(0)>>1) - 1

	var newTotalTrips float64
	for left < right {
		mid := left + (right-left)/2
		newTotalTrips = 0
		for j := range time {
			newTotalTrips += float64(mid) / float64(time[j])
		}
		if newTotalTrips == float64(totalTrips) {
			return int64(left)
		} else if int(newTotalTrips) < totalTrips {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int64(left + 1)
}

func main() {
	time := []int{5, 10, 10}
	totalTrips := 9
	fmt.Println(minimumTime(time, totalTrips))
	fmt.Println(minimumTime2(time, totalTrips))
}
