package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var n int
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}

	N := k
	if n < k {
		N = k % n
	}

	newHead := head
	for i := 0; i < N; i++ {
		newHead = rotate(head)
		head = newHead
	}
	return newHead
}

func rotate(head *ListNode) *ListNode {

	cur := head
	next := cur.Next
	for next.Next != nil {
		cur = cur.Next
		next = cur.Next
	}
	cur.Next = nil
	next.Next = head

	return next
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	cur := head
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head.Next
}

func (lst *ListNode) print() {
	cur := lst
	for cur != nil {
		if cur.Next == nil {
			fmt.Println(cur.Val)
		} else {
			fmt.Print(cur.Val, " -> ")
		}
		cur = cur.Next
	}
}

func main() {
	head := []int{1, 2, 3, 4, 5}
	k := 6
	lst := createList(head)
	lst.print()
	lst2 := rotateRight(lst, k)
	lst2.print()
}
