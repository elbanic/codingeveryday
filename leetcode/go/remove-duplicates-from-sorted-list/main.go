package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func createList(nums []int) *ListNode {

	head := &ListNode{}
	cur := head
	for _, n := range nums {
		cur.Next = &ListNode{Val: n}
		cur = cur.Next
	}
	return head.Next
}

func (list *ListNode) print() {
	cur := list
	for cur != nil {
		if cur.Next != nil {
			fmt.Print(cur.Val, " -> ")
		} else {
			fmt.Println(cur.Val)
		}
		cur = cur.Next
	}
}

func main() {

	head := []int{1, 1, 1}
	list := createList(head)
	fmt.Println(deleteDuplicates(list))
	list.print()
}
