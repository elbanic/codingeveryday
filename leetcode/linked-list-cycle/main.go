package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	visited := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if visited[cur] {
			return true
		}
		visited[cur] = true
		cur = cur.Next
	}
	return false
}

func createLinkedList(nums []int, pos int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, v := range nums {
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
	}

	if pos < 0 {
		return head.Next
	}
	last := head.Next
	for i := 0; i < pos; i++ {
		last = last.Next
	}
	cur.Next = last
	return head.Next
}

func main() {
	head := createLinkedList([]int{3, 2, 0, -4}, 1)
	fmt.Println(hasCycle(head))
}
