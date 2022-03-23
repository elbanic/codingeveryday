/*
https://leetcode.com/problems/middle-of-the-linked-list/
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

Example 1:
	Input: head = [1,2,3,4,5]
	Output: [3,4,5]
	Explanation: The middle node of the list is node 3.

Example 2:
	Input: head = [1,2,3,4,5,6]
	Output: [4,5,6]
	Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var size int
	for cur != nil {
		cur = cur.Next
		size++
	}

	cur = head
	for i := 0; i < size/2; i++ {
		cur = cur.Next
	}
	return cur
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

func (list *ListNode) print() {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	head := createLinkedList(nums)
	head.print()
	middle := middleNode(head)
	middle.print()
}
