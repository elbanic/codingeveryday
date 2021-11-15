/*
https://leetcode.com/problems/kth-largest-element-in-a-stream/
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.
Implement KthLargest class:
* KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
* int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.

Example 1:

	Input
	["KthLargest", "add", "add", "add", "add", "add"]
	[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
	Output
	[null, 4, 5, 5, 8, 8]

	Explanation
	KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
	kthLargest.add(3);   // return 4
	kthLargest.add(5);   // return 5
	kthLargest.add(10);  // return 5
	kthLargest.add(9);   // return 8
	kthLargest.add(4);   // return 8
 */

package main

import (
	"container/heap"
	"fmt"
)

type maxIntHeap []int

func (h maxIntHeap) Len() int           { return len(h) }
func (h maxIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h maxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxIntHeap) Peak() int{
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

type KthLargest struct {
	maxHeap *maxIntHeap
	kth     int
}

func Constructor(k int, nums []int) KthLargest {
	h := &maxIntHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}
	kthLargest := KthLargest{h, k}
	for kthLargest.maxHeap.Len() > k {
		heap.Pop(kthLargest.maxHeap)
	}
	return kthLargest
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k.maxHeap, val)
	if k.maxHeap.Len() > k.kth {
		heap.Pop(k.maxHeap)
	}
	return k.maxHeap.Peak()
}

func (k KthLargest) peak() int {
	return k.maxHeap.Peak()
}

func main() {

	k, nums := 3, []int{4, 5, 8, 2}
	k1 := Constructor(k, nums)

	fmt.Println(k1.Add(3))
	fmt.Println(k1.Add(5))
	fmt.Println(k1.Add(10))
	fmt.Println(k1.Add(9))
	fmt.Println(k1.Add(4))
}
