package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//BFS
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := [][]*Node{[]*Node{root}}
	for len(queue) > 0 {
		curLevel := queue[0]
		queue = queue[1:]
		var nextLevel []*Node
		for i := range curLevel {
			if i < len(curLevel)-1 {
				curLevel[i].Next = curLevel[i+1]
			}
			if curLevel[i].Left != nil {
				nextLevel = append(nextLevel, curLevel[i].Left)
			}
			if curLevel[i].Right != nil {
				nextLevel = append(nextLevel, curLevel[i].Right)
			}
		}
		if nextLevel != nil {
			queue = append(queue, nextLevel)
		}
	}
	return root
}

func createBTree(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	root := Node{nums[0], nil, nil, nil}
	queue := []*Node{&root}
	for i := 1; i < len(nums)-1; i += 2 {
		cur := queue[0]
		queue = queue[1:]

		left := &Node{nums[i], nil, nil, nil}
		cur.Left = left
		queue = append(queue, left)

		right := &Node{nums[i+1], nil, nil, nil}
		cur.Right = right
		queue = append(queue, right)
	}
	return &root
}

func main() {
	root := []int{1, 2, 3, 4, 5, 6, 7}
	btree := createBTree(root)
	connected := connect(btree)
	fmt.Println(connected)
}
