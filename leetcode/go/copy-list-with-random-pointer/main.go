package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	var nodeArr []*Node
	randomIndex := make(map[int]int)

	cur := head
	for cur != nil {
		nodeArr = append(nodeArr, cur)
		cur = cur.Next
	}

	for i, node := range nodeArr {
		for j, node2 := range nodeArr {
			if node2 == node.Random {
				randomIndex[i] = j
			}
		}
	}

	var nodeArr2 []*Node
	newHaed := &Node{}
	cur = newHaed
	for _, node := range nodeArr {
		cur.Next = &Node{Val: node.Val}
		nodeArr2 = append(nodeArr2, cur.Next)
		cur = cur.Next
	}

	cur = newHaed.Next
	for i := range nodeArr2 {
		if id, ok := randomIndex[i]; ok {
			cur.Random = nodeArr2[id]
		} else {
			cur.Random = nil
		}
		cur = cur.Next
	}
	return newHaed.Next
}

func createListNode(nums [][]int) *Node {
	if len(nums) == 0 {
		return nil
	}
	var arr []*Node
	head := &Node{}
	cur := head
	for _, n := range nums {
		cur.Next = &Node{Val: n[0]}
		arr = append(arr, cur.Next)
		cur = cur.Next
	}

	cur = head.Next
	for _, n := range nums {
		if n[1] != -1 {
			cur.Random = arr[n[1]]
		}
		cur = cur.Next
	}
	return head.Next
}

func (head *Node) print() {
	cur := head
	for cur != nil {
		if cur.Next == nil {
			if cur.Random != nil {
				fmt.Println("[", cur.Val, ",", cur.Random.Val, "]")
			} else {
				fmt.Println("[", cur.Val, ",nil]")
			}

		} else {
			if cur.Random != nil {
				fmt.Print("[", cur.Val, ",", cur.Random.Val, "] -> ")
			} else {
				fmt.Print("[", cur.Val, ",nil] -> ")
			}
		}
		cur = cur.Next
	}
}

func main() {
	head := [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	nodes := createListNode(head)
	nodes.print()

	nodes2 := copyRandomList(nodes)
	nodes2.print()
}
