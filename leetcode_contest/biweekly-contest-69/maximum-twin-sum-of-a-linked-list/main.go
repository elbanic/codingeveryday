package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var arr []*ListNode

	curr := head
	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}

	preTail := &ListNode{}
	curr = preTail
	for i := len(arr) - 1; i >= 0; i-- {
		curr.Next = &ListNode{arr[i].Val, nil}
		curr = curr.Next
	}
	currTail := preTail.Next
	currHead := head
	var max int
	for i := 0; i < len(arr)/2; i++ {
		max = int(math.Max(float64(max), float64(currTail.Val+currHead.Val)))
		currTail = currTail.Next
		currHead = currHead.Next
	}
	return max
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := ListNode{Val: nums[0], Next: nil}
	cur := &head
	for i := 1; i < len(nums); i++ {
		next := ListNode{Val: nums[i], Next: nil}
		cur.Next = &next
		cur = cur.Next
	}
	return &head
}

func main() {
	head := []int{5, 4, 2, 1}
	fmt.Println(pairSum(createLinkedList(head)))
}
