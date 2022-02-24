package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//divide and conquer
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	left := sortList2(head)
	right := sortList2(mid)

	return merge(left, right)
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	tail := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
			tail = tail.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
			tail = tail.Next
		}
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummyHead.Next
}

func getMid(head *ListNode) *ListNode {
	var midPrev *ListNode
	for head != nil && head.Next != nil {
		if midPrev == nil {
			midPrev = head
		} else {
			midPrev = midPrev.Next
		}
		head = head.Next.Next
	}
	mid := midPrev.Next
	midPrev.Next = nil
	return mid
}

func sortList(head *ListNode) *ListNode {

	var arr []*ListNode
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	newHead := &ListNode{}
	cur = newHead
	for i := range arr {
		cur.Next = arr[i]
		cur = cur.Next
	}
	cur.Next = nil
	return newHead.Next
}

func createList(nums []int) *ListNode {

	head := &ListNode{}
	cur := head
	for i := range nums {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head.Next
}

func (lst *ListNode) print() {
	cur := lst
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	head := []int{4, 2, 1, 3}
	lst := createList(head)
	lst.print()
	sorted := sortList2(lst)
	sorted.print()
}
