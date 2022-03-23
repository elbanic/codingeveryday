/*
https://leetcode.com/problems/merge-two-sorted-lists/
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
	Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]

Example 2:
	Input: list1 = [], list2 = []
	Output: []

Example 3:
	Input: list1 = [], list2 = [0]
	Output: [0]
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//my answer
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	head := &ListNode{}
	cur := head
	cur1 := list1
	cur2 := list2
	for cur1 != nil || cur2 != nil {
		var val int
		if cur1 != nil && cur2 != nil {
			if cur1.Val < cur2.Val {
				val = cur1.Val
				cur1 = cur1.Next
			} else {
				val = cur2.Val
				cur2 = cur2.Next
			}
		} else {
			if cur1 != nil {
				val = cur1.Val
				cur1 = cur1.Next
			} else {
				val = cur2.Val
				cur2 = cur2.Next
			}
		}
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return head.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	head := &ListNode{}
	cur := head
	cur1 := list1
	cur2 := list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	if cur1 == nil {
		cur.Next = cur2
	} else {
		cur.Next = cur1
	}
	return head.Next
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{}
	cur := head
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return head.Next
}

func printList(list *ListNode) {
	if list == nil {
		return
	}

	cur := list
	for cur != nil {
		if cur.Next != nil {
			fmt.Printf("%d -> ", cur.Val)
		} else {
			fmt.Printf("%d\n", cur.Val)
		}
		cur = cur.Next
	}
}

func main() {
	list1 := createList([]int{5})
	list2 := createList([]int{1, 2, 4})
	printList(mergeTwoLists(list1, list2))
}
