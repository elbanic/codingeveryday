package main

import (
	"container/heap"
	"fmt"
	"math"
)

type maxIntHeap []int

func (h maxIntHeap) Len() int           { return len(h) }
func (h maxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxIntHeap) Peak() int {
	return h[0]
}

func (h *maxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumTop(nums []int, k int) int {
	maxHeap := &maxIntHeap{}
	heap.Init(maxHeap)
	output := helper(nums, maxHeap, k, -1)

	return output
}

func helper(nums []int, removed *maxIntHeap, k, max int) int {
	if k == 0 {
		if len(nums) == 0 {
			return -1
		} else {
			return nums[0]
		}
	}

	if removed.Len() == 0 {
		heap.Push(removed, nums[0])
		nums = nums[1:]
		max = helper(nums, removed, k-1, max)
	} else if len(nums) == 0 {
		nums = append(nums, removed.Peak())
		heap.Pop(removed)
		max = helper(nums, removed, k-1, max)
	} else {
		cpRemoved1 := make(maxIntHeap, removed.Len())
		copy(cpRemoved1, *removed)
		heap.Push(&cpRemoved1, nums[0])
		cpNums1 := make([]int, len(nums))
		copy(cpNums1, nums)
		cpNums1 = cpNums1[1:]

		cpNums2 := make([]int, len(nums))
		copy(cpNums2, nums)
		cpNums2 = append([]int{removed.Peak()}, cpNums2...)
		cpRemoved2 := make(maxIntHeap, removed.Len())
		copy(cpRemoved2, *removed)
		heap.Pop(&cpRemoved2)

		max = int(math.Max(float64(helper(cpNums2, &cpRemoved2, k-1, max)), float64(helper(cpNums1, &cpRemoved1, k-1, max))))
	}

	return max
}

func main() {
	nums := []int{99, 95, 68, 24, 18}
	k := 30
	fmt.Println(maximumTop(nums, k))
}
