package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	tail := head.Next
	temp := tail.Next
	tail.Next = head
	head.Next = swapPairs(temp)

	return tail
}

func helper(head *ListNode, tail *ListNode) *ListNode {
	head.Next = tail.Next
	tail.Next = head
	return tail
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

func main() {
	head := []int{1}
	lst := createList(head)
	lst = swapPairs(lst)
	fmt.Println(lst)
}
