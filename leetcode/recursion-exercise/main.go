package main

import "fmt"

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString1(s1)
	fmt.Printf("%s\n", s1)

	s2 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString2(s2)
	fmt.Printf("%s\n", s2)

	n4 := ListNode{Val: 4}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	printNode(swapPairs(&n1))
}

func reverseString1(s []byte) {
	reverseStringDivide(s, 0, len(s)-1)
}

//Divide-and-Conquer solution, O(1) space complexity
func reverseStringDivide(s []byte, start, end int) {
	if start >= end {
		return
	}
	s[start], s[end] = s[end], s[start]
	reverseStringDivide(s, start+1, end-1)
}

func reverseString2(s []byte) {
	newS := reverseStringHelper(s, nil)
	for i := 0; i < len(s); i++ {
		s[i] = newS[i]
	}
}

//recursion, O(N) space complexity
func reverseStringHelper(s []byte, output []byte) []byte {
	if len(s) <= 0 {
		return nil
	}
	c := s[0]
	output = reverseStringHelper(s[1:], output)
	output = append(output, c)
	return output
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
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