package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {

	var stack []*TreeNode
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)

	for len(stack) > 0 {
		var curLevel []*TreeNode
		for i := len(stack) - 1; i >= 0; i-- {
			if u == stack[i] {
				if len(curLevel) > 0 {
					return curLevel[len(curLevel)-1]
				} else {
					return nil
				}
			}
			if stack[i] != nil {
				curLevel = append(curLevel, stack[i])
			}
		}

		stack = []*TreeNode{}
		for i := len(curLevel) - 1; i >= 0; i-- {
			stack = append(stack, curLevel[i].Left)
			stack = append(stack, curLevel[i].Right)
		}
	}
	return nil
}

func createTree(nums []int, u int) (*TreeNode, *TreeNode) {
	root := &TreeNode{Val: nums[0]}
	var queue []*TreeNode
	queue = append(queue, root)

	var uPointer *TreeNode
	for i := 1; i < len(nums); i += 2 {
		cur := queue[0]
		queue = queue[1:]

		cur.Left = &TreeNode{Val: nums[i]}
		queue = append(queue, cur.Left)
		if nums[i] == u {
			uPointer = cur.Left
		}

		if i+1 < len(nums) {
			cur.Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, cur.Right)
			if nums[i+1] == u {
				uPointer = cur.Right
			}
		}
	}
	return root, uPointer
}

func (tr *TreeNode) printTree() {
	fmt.Print(tr.Val)
	if tr.Left != nil {
		tr.Left.printTree()
	}
	if tr.Right != nil {
		tr.Right.printTree()
	}
}

func main() {

	tr, u := createTree([]int{1, 2, 3, -1, 4, 5, 6}, 4)
	fmt.Println(findNearestRightNode(tr, u).Val)

	tr2, u2 := createTree([]int{3, -1, 4, 2}, 2)
	fmt.Println(findNearestRightNode(tr2, u2))

}
