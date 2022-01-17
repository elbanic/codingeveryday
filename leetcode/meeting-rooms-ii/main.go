/*
https://leetcode.com/problems/meeting-rooms-ii/
Given an array of meeting time intervals intervals where intervals[i] = [starti, endi],
return the minimum number of conference rooms required.

Example 1:
	Input: intervals = [[0,30],[5,10],[15,20]]
	Output: 2

Example 2:
	Input: intervals = [[7,10],[2,4]]
	Output: 1
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//priority queue
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Peak() int {
	return h[0]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms2(intervals [][]int) int {
	minH := &minHeap{}
	heap.Init(minH)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	heap.Push(minH, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= minH.Peak() {
			heap.Pop(minH)
		}
		heap.Push(minH, intervals[i][1])
	}
	return minH.Len()
}

//brute force
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var rooms [][]int
	rooms = append(rooms, []int{0})
	roomIdx := 1

	for roomIdx < len(intervals) {
		cur := intervals[roomIdx]
		var overlapped bool
		for i, ids := range rooms {
			overlapped = false
			for _, id := range ids {
				if isOverlapped(intervals[id], cur) {
					overlapped = true
					break
				}
			}
			if !overlapped {
				rooms[i] = append(rooms[i], roomIdx)
				break
			}
		}
		if overlapped {
			rooms = append(rooms, []int{roomIdx})
		}
		roomIdx++
	}
	return len(rooms)
}

func isOverlapped(a, b []int) bool {
	if a[0] <= b[0] && a[1] > b[0] {
		return true
	}
	return false
}

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(minMeetingRooms2(intervals))

	intervals2 := [][]int{{7, 10}, {2, 4}}
	fmt.Println(minMeetingRooms2(intervals2))

	intervals3 := [][]int{{1, 10}, {2, 7}, {3, 19}, {8, 12}, {10, 20}, {11, 30}}
	fmt.Println(minMeetingRooms2(intervals3))
}
