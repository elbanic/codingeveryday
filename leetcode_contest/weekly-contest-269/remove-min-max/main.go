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

	leftDelId := int(math.Min(float64(maxId), float64(minId)))
	rightDelId := int(math.Max(float64(maxId), float64(minId)))
	if leftDelId >= len(nums)/2 {
		return len(nums) - leftDelId
	} else if rightDelId <= len(nums)/2 {
		return rightDelId + 1
	}

	leftGap := leftDelId + 1
	rightGap := len(nums) - rightDelId
	if leftGap < rightGap {
		if rightDelId - leftDelId < rightGap {
			return rightDelId + 1
		} else {
			return rightGap + leftGap
		}
	} else {
		if -(leftDelId - rightDelId) < leftGap {
			return len(nums) - leftDelId
		} else {
			return rightGap + leftGap
		}
	}
}

func main() {

	nums := []int{41,-47,-26,57,82,-23,47,52,42,79,2,77,0,-4,1,-99,-57,72,-95}
	fmt.Println(minimumDeletions(nums))

}
