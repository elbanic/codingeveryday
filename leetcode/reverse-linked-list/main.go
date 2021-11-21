/*
https://leetcode.com/problems/reverse-linked-list/
Given the head of a singly linked list, reverse the list, and return the reversed list.
Example 1:
	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

Example 2:
	Input: head = [1,2]
	Output: [2,1]
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return prev
}

func main() {
	n5 := ListNode{Val: 5}
	n4 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	printNode(reverseList(&n1))
}

func printNode(head *ListNode) {
	cur := head
	for cur != nil{
		if cur.Next == nil {
			fmt.Printf("%d\n", cur.Val)
		} else {
			fmt.Printf("%d -> ", cur.Val)
		}
		cur = cur.Next
	}
}

