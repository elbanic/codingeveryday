package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//tortoise and hare
func detectCycle2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	intersect := getIntersect(head)
	if intersect == nil {
		return nil
	}

	ptr1 := head
	ptr2 := intersect
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	if ptr1 == ptr2 {
		return ptr1
	}
	return nil
}

func getIntersect(head *ListNode) *ListNode {
	if head == head.Next {
		return head
	}
	fast := head.Next
	var slow *ListNode
	if head.Next != nil {
		slow = head.Next.Next
	}

	for fast != slow {
		if fast == nil || slow == nil {
			return nil
		}
		fast = fast.Next
		if slow.Next != nil {
			slow = slow.Next.Next
		} else {
			return nil
		}
	}
	return fast
}

//hash table
func detectCycle(head *ListNode) *ListNode {

	m := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if m[cur] {
			return cur
		}
		m[cur] = true
		cur = cur.Next
	}
	return nil
}

func createLinkedListWithCycle(nums []int, pos int) *ListNode {

	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}

	//make cycle until pos
	cycle := head
	for i := 0; i < pos; i++ {
		cycle = cycle.Next
	}
	cur.Next = cycle
	return head
}

func main() {
	nums := []int{3, 2, 0, -4}
	pos := 1
	head := createLinkedListWithCycle(nums, pos)
	fmt.Println(detectCycle2(head).Val)
}
