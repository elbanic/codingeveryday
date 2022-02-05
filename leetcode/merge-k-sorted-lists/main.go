package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Heap []*ListNode

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Heap) Swap(i, j int) {
	h[i].Val, h[j].Val = h[j].Val, h[i].Val
}

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists2(lists []*ListNode) *ListNode {

	h := &Heap{}
	heap.Init(h)

	for i := 0; i < len(lists); i++ {
		cur := lists[i]
		for cur != nil {
			heap.Push(h, cur)
			cur = cur.Next
		}
	}

	head := &ListNode{}
	cur := head
	for h.Len() > 0 {
		cur.Next = &ListNode{Val: heap.Pop(h).(*ListNode).Val}
		cur = cur.Next
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {

	var arr []int
	for i := 0; i < len(lists); i++ {
		cur := lists[i]
		for cur != nil {
			arr = append(arr, cur.Val)
			cur = cur.Next
		}
	}

	sort.Ints(arr)
	return createList(arr)
}

func createList(nums []int) *ListNode {

	head := &ListNode{}
	cur := head
	for _, i := range nums {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}

	return head.Next
}

func createLinkedList(nums [][]int) []*ListNode {
	var ret []*ListNode

	for _, i := range nums {
		head := &ListNode{}
		cur := head
		for _, j := range i {
			cur.Next = &ListNode{Val: j}
			cur = cur.Next
		}
		ret = append(ret, head.Next)
	}
	return ret
}

func (list *ListNode) print() {
	cur := list
	for cur != nil {
		fmt.Print(cur.Val, " -> ")
		cur = cur.Next
	}
}

func main() {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	l := createLinkedList(lists)

	mergeKLists2(l).print()

}
