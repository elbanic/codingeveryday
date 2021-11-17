/*
https://leetcode.com/problems/top-k-frequent-elements/
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

Example 2:
	Input: nums = [1], k = 1
	Output: [1]
 */

package main

import (
	"fmt"
)

type node struct {
	key int
	value int
}

type maxHeap struct {
	maximumHeap []node
	heapSize    int
	realSize    int
}

func newMaxHeap(heapSize int) maxHeap {
	return maxHeap{maximumHeap: []node{node{0,0}}, heapSize: heapSize+1, realSize: 0}
}

func (h *maxHeap) add(n node) {
	h.realSize++
	if h.realSize > h.heapSize {
		fmt.Println("Add too many elements!")
		h.realSize--
		return
	}

	h.maximumHeap = append(h.maximumHeap, n)
	index := h.realSize
	parent := index / 2

	for h.maximumHeap[index].value > h.maximumHeap[parent].value && index > 1 {
		temp := h.maximumHeap[index]
		h.maximumHeap[index] = h.maximumHeap[parent]
		h.maximumHeap[parent] = temp
		index = parent
		parent = index / 2
	}
}

func (h maxHeap) peak() node{
	return h.maximumHeap[1]
}

func (h *maxHeap) pop() node{

	if h.realSize < 1 {
		fmt.Println("Don't have any element!")
		return node{}
	} else {
		removeElement := h.maximumHeap[1]
		h.maximumHeap[1] = h.maximumHeap[h.realSize]
		h.realSize--
		index := 1

		for index < h.realSize && index <= h.realSize / 2 {
			left := index * 2
			right := (index * 2) + 1

			if h.maximumHeap[index].value < h.maximumHeap[left].value ||
				h.maximumHeap[index].value < h.maximumHeap[right].value {
				if h.maximumHeap[left].value > h.maximumHeap[right].value {
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
		fmt.Print(h.maximumHeap[h.realSize])
		fmt.Println("]")
	}
}

func topKFrequent(nums []int, k int) []int {

	mapNums := make(map[int]int)
	for _,v := range nums {
		if _, ok := mapNums[v]; ok {
			mapNums[v] += 1
		} else {
			mapNums[v] = 1
		}
	}
	maxH := newMaxHeap(len(mapNums))
	for k,v := range mapNums {
		maxH.add(node{k, v})
	}

	var ret []int
	for i:=0; i<k; i++ {
		node := maxH.pop()
		ret = append(ret, node.key)
	}
	return ret
}

func main() {

	nums1, k := []int{1,1,1,2,2,3}, 2
	fmt.Println(topKFrequent(nums1, k))

	nums2, k := []int{1}, 1
	fmt.Println(topKFrequent(nums2, k))

	nums3, k := []int{-1, -1}, 1
	fmt.Println(topKFrequent(nums3, k))

	nums4, k := []int{1,1,1,2,2,2,3,3,4,5,5,5,5}, 3
	fmt.Println(topKFrequent(nums4, k))
}

