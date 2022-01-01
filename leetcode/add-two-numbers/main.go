/*
https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
	Input: l1 = [2,4,3], l2 = [5,6,4]
	Output: [7,0,8]
	Explanation: 342 + 465 = 807.

Example 2:
	Input: l1 = [0], l2 = [0]
	Output: [0]

Example 3:
	Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	Output: [8,9,9,9,0,0,0,1]
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//iteration
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{}
	cur, cur1, cur2 := head, l1, l2
	var carry int
	for cur1 != nil || cur2 != nil {
		var val1, val2 int
		if cur1 != nil {
			val1 = cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			val2 = cur2.Val
			cur2 = cur2.Next
		}
		sum := val1 + val2 + carry
		carry = sum / 10

		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return head.Next
}

//recursion
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	helper(l1, l2, head, 0)
	return head
}

func helper(l1 *ListNode, l2 *ListNode, add *ListNode, extra int) {
	if l1 == nil && l2 == nil {
		return
	}
	var l1val, l2val int
	var l1next, l2next *ListNode
	if l1 == nil {
		l2val = l2.Val
		l2next = l2.Next
	} else if l2 == nil {
		l1val = l1.Val
		l1next = l1.Next
	} else {
		l1val = l1.Val
		l2val = l2.Val
		l1next = l1.Next
		l2next = l2.Next
	}
	add.Val = (l1val + l2val + extra) % 10
	if l1val+l2val+extra >= 10 {
		extra = 1
	} else {
		extra = 0
	}
	if l1next == nil && l2next == nil {
		if extra == 0 {
			add.Next = nil
		} else {
			add.Next = &ListNode{1, nil}
		}
	} else {
		add.Next = &ListNode{}
	}
	helper(l1next, l2next, add.Next, extra)
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

func printLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		cur = cur.Next
		fmt.Println()
	}
}

func main() {
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4})
	res := addTwoNumbers2(l1, l2)
	printLinkedList(res)
}
