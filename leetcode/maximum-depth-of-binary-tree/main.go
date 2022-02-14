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

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth = int(math.Max(float64(dfs(root.Left, depth+1)), float64(dfs(root.Right, depth+1))))
	return depth
}

func createTree(nums []int) *TreeNode {

	head := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{head}

	for i := 1; i < len(nums); i += 2 {
		cur := queue[0]
		queue = queue[1:]

		if i < len(nums) {
			if nums[i] != -1 {
				cur.Left = &TreeNode{Val: nums[i]}
				queue = append(queue, cur.Left)
			}
		}
		if i+1 < len(nums) {
			if nums[i+1] != -1 {
				cur.Right = &TreeNode{Val: nums[i+1]}
				queue = append(queue, cur.Right)
			}
		}
	}
	return head
}

func main() {
	root1 := []int{3, 9, 20, -1, -1, 15, 7}
	r1 := createTree(root1)
	fmt.Println(maxDepth(r1))

	root2 := []int{1, -1, 2}
	r2 := createTree(root2)
	fmt.Println(maxDepth(r2))
}
