package main

import (
	"container/heap"
	"fmt"
	"math"
)

type valueId struct {
	val int
	id int
}

type maxHeap []valueId

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, valueId{x.(valueId).val, x.(valueId).id})
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type minHeap []valueId

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, valueId{x.(valueId).val, x.(valueId).id})
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumDeletions(nums []int) int {
	maxh := &maxHeap{}
	heap.Init(maxh)

	minh := &minHeap{}
	heap.Init(minh)

	for i,v := range nums {
		newVal := valueId{v, i}
		heap.Push(maxh, newVal)
		heap.Push(minh, newVal)
	}

	maxId := heap.Pop(maxh).(valueId).id
	minId := heap.Pop(minh).(valueId).id

	if maxId == minId {
		return 1
	}

	var leftDelId, rightDelId int
	if maxId >= len(nums)/2 {
		rightDelId = len(nums) - maxId
	} else {
		leftDelId = maxId
	}

	if minId >= len(nums)/2 {
		rightDelId = int(math.Max(float64(rightDelId), float64(len(nums) - minId)))
	} else {
		leftDelId = int(math.Max(float64(leftDelId), float64(minId)))
	}

	return rightDelId + (leftDelId + 1)
}

func main() {

	nums := []int{101}
	fmt.Println(minimumDeletions(nums))

}
