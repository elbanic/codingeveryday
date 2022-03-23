/*
Given a singly linked list, return a random node's value from the linked list.
Each node must have the same probability of being chosen.
Implement the Solution class:
	* Solution(ListNode head) Initializes the object with the integer array nums.
	* int getRandom() Chooses a node randomly from the list and returns its value.
		All the nodes of the list should be equally likely to be choosen.

Example 1:
	Input
	["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
	[[[1, 2, 3]], [], [], [], [], []]
	Output
	[null, 1, 3, 2, 2, 3]

	Explanation
	Solution solution = new Solution([1, 2, 3]);
	solution.getRandom(); // return 1
	solution.getRandom(); // return 3
	solution.getRandom(); // return 2
	solution.getRandom(); // return 2
	solution.getRandom(); // return 3
	// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	size int
}

func Constructor(head *ListNode) Solution {
	cur := head
	var size int
	for cur != nil {
		size++
		cur = cur.Next
	}
	return Solution{head, size}
}

func (this *Solution) GetRandom() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randInt := r1.Intn(this.size)

	cur := this.head
	for i := 0; i < randInt; i++ {
		cur = cur.Next
	}
	return cur.Val
}

//Reservoir Sampleing
func (this *Solution) GetRandom2() int {
	scope := 1
	chosenValue := 0
	cur := this.head

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for cur != nil {
		if r1.Float64() < 1.0/float64(scope) {
			chosenValue = cur.Val
		}
		scope += 1
		cur = cur.Next
	}
	return chosenValue
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
	return head.Next
}

func main() {
	sol := Constructor(createLinkedList([]int{1, 2, 3}))
	fmt.Println(sol.GetRandom2())
	fmt.Println(sol.GetRandom2())
	fmt.Println(sol.GetRandom2())
	fmt.Println(sol.GetRandom2())
	fmt.Println(sol.GetRandom2())
}
