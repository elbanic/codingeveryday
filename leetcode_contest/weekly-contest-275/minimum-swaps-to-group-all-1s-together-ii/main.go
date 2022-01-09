package main

import "fmt"

func minSwaps(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	head := createLinkedList(nums)
	cur := head

	var ret int
	for cur != nil && cur.Next != head {
		if cur.Val == 1 {
			if isValid(cur) {
				return ret
			}
		}
		cur = cur.Next
	}
	return 0
}

func isValid(head *ListNode) bool {

	cur := head
	cur = cur.Next
	var stat bool
	if cur.Val == 1 {
		stat = true
	}

	for cur != nil && cur != head {
		if cur.Val == 1 && !stat {
			return false
		}
		if cur.Val == 0 {
			stat = false
		}
		cur = cur.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	cur := head
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	cur.Next = head.Next
	return head.Next
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 0, 1, 1, 0}
	fmt.Println(minSwaps(nums))
}
