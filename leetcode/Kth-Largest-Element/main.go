/*
https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/
Given an integer array `nums` and an integer `k`, return *the* `k th` *largest element in the array*.
Note that it is the `K th` largest element in the sorted order, not the `k th` distinct element.

Example 1:
	Input: nums = [3,2,1,5,6,4], k = 2
	Output: 5

Example 2:
	Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
	Output: 4

 */

package main

import (
	"fmt"
	"strconv"
)

type maxHeap struct {
	maximumHeap []int
	heapSize    int
	realSize    int
}

func newMaxHeap(heapSize int) maxHeap {
	return maxHeap{maximumHeap: []int{0}, heapSize: heapSize+1, realSize: 0}
}

func (h *maxHeap) add(e int) {
	h.realSize++
	if h.realSize > h.heapSize {
		fmt.Println("Add too many elements!")
		h.realSize--
		return
	}

	h.maximumHeap = append(h.maximumHeap, e)
	index := h.realSize
	parent := index / 2

	for h.maximumHeap[index] > h.maximumHeap[parent] && index > 1 {
		temp := h.maximumHeap[index]
		h.maximumHeap[index] = h.maximumHeap[parent]
		h.maximumHeap[parent] = temp
		index = parent
		parent = index / 2
	}
}

func (h maxHeap) peak() int{
	return h.maximumHeap[1]
}

func (h *maxHeap) pop() int{

	if h.realSize < 1 {
		fmt.Println("Don't have any element!")
		return int(^uint(0) >> 1) // integer MAX
	} else {
		removeElement := h.maximumHeap[1]
		h.maximumHeap[1] = h.maximumHeap[h.realSize]
		h.realSize--
		index := 1

		for index < h.realSize && index <= h.realSize / 2 {
			left := index * 2
			right := (index * 2) + 1

			if h.maximumHeap[index] < h.maximumHeap[left] || h.maximumHeap[index] < h.maximumHeap[right] {
				if h.maximumHeap[left] > h.maximumHeap[right] {
					temp := h.maximumHeap[left]
					h.maximumHeap[left] = h.maximumHeap[index]
					h.maximumHeap[index] = temp
					index = left
				} else {
					temp := h.maximumHeap[right]
					h.maximumHeap[right] = h.maximumHeap[index]
					h.maximumHeap[index] = temp
					index = right
				}
			} else {
				break
			}
		}
		h.maximumHeap = h.maximumHeap[0: h.realSize+1]
		return removeElement
	}
}

func (h maxHeap) size() int{
	return h.realSize
}

func (h maxHeap) print() {
	if h.realSize == 0 {
		fmt.Println("No element")
	} else {
		fmt.Print("[")
		for i := 1; i<= h.realSize-1; i++ {
			fmt.Print(h.maximumHeap[i], ",")
		}
		fmt.Println(strconv.Itoa(h.maximumHeap[h.realSize]) + "]")
	}
}

func findKthLargest(nums []int, k int) int {
	maxH := newMaxHeap(len(nums))
	for _,v := range nums {
		maxH.add(v)
	}
	for i:=1; i<k; i++{
		maxH.pop()
	}
	return maxH.peak()
}

func main() {
	nums1, k := []int{3,2,1,5,6,4}, 2
	fmt.Println(findKthLargest(nums1, k))

	nums2, k := []int{3,2,3,1,2,4,5,5,6}, 4
	fmt.Println(findKthLargest(nums2, k))
}
