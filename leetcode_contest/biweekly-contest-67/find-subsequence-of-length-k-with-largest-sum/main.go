
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type orderInt struct {
	val int
	order int
}

type IntHeap []orderInt

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h IntHeap) Swap(i, j int)      {
	h[i].val, h[j].val = h[j].val, h[i].val
	h[i].order, h[j].order = h[j].order, h[i].order
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(orderInt))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSubsequence(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)
	for i,v := range nums {
		heap.Push(h, orderInt{v, i})
	}
	var res []orderInt
	for k > 0 {
		top := heap.Pop(h).(orderInt)
		res = append(res, top)
		k--
	}

	var ret []int
	sort.Slice(res, func(i, j int) bool {
		return res[i].order < res[j].order
	})
	for _,v := range res{
		ret = append(ret, v.val)
	}
	return ret
}

func main(){
	nums := []int{-1,-2,3,4}
	k := 3

	fmt.Println(maxSubsequence(nums, k))
}
