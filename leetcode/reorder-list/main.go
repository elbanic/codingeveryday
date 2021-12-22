/*
https://leetcode.com/problems/reorder-list/
You are given the head of a singly linked-list. The list can be represented as:

	L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example1:
	Input: head = [1,2,3,4]
	Output: [1,4,2,3]

Example2:
	Input: head = [1,2,3,4,5]
	Output: [1,5,2,4,3]
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	var arr []*ListNode
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}

	cur = head
	for i := len(arr) - 1; i >= 0; i-- {
		if cur != arr[i] {
			next := cur.Next
			cur.Next = arr[i]
			arr[i].Next = next
			if arr[i] == next || arr[i] == next.Next {
				next.Next = nil
				break
			} else {
				cur = next
			}
		} else {
			break
		}
	}
}

func main() {
	//e := ListNode{5, nil}
	a := ListNode{1, nil}
	b := ListNode{1, &a}
	c := ListNode{1, &b}
	d := ListNode{1, &c}

	reorderList(&d)
	fmt.Println("")
}
