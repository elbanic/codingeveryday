package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {

	if head != nil && head.Val == 0 {
		var sum int
		cur := head.Next
		if cur == nil {
			return cur
		}
		for cur.Val != 0 {
			sum += cur.Val
			cur = cur.Next
		}
		head = &ListNode{Val: sum}
		head.Next = mergeNodes(cur)
	}
	return head
}

func createLinkedList(nums []int) *ListNode {

	head := &ListNode{}
	cur := head
	for i := range nums {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head.Next
}

func (lst *ListNode) printList() {
	cur := lst
	for cur != nil {
		fmt.Print(cur.Val, " -> ")
		cur = cur.Next
	}
	fmt.Println()
}

func main() {

	head := []int{0, 3, 1, 0, 4, 5, 2, 0}
	headlst := createLinkedList(head)
	mergeNodes(headlst).printList()
}
