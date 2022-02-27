package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type pair struct {
		node *TreeNode
		id   int
	}
	var maxW int
	queue := []*pair{&pair{root, 0}}

	for len(queue) > 0 {
		head := queue[0]
		curLevelSize := len(queue)
		var elem *pair
		for i := 0; i < curLevelSize; i++ {
			elem = queue[0]
			queue = queue[1:]
			if elem.node.Left != nil {
				queue = append(queue, &pair{elem.node.Left, 2 * elem.id})
			}
			if elem.node.Right != nil {
				queue = append(queue, &pair{elem.node.Right, 2*elem.id + 1})
			}
		}
		maxW = int(math.Max(float64(maxW), float64(elem.id-head.id+1)))
	}
	return maxW
}

func createTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(nums); i += 2 {
		cur := queue[0]
		queue = queue[1:]
		if nums[i] != -1 {
			cur.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, cur.Left)
		}
		if i+1 < len(nums) && nums[i+1] != -1 {
			cur.Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, cur.Right)
		}
	}
	return root
}

func main() {

	root := []int{1, 3, 2, 5, -1, -1, 9, 6, -1, -1, 7}
	tree := createTree(root)
	fmt.Println(widthOfBinaryTree(tree))
}
