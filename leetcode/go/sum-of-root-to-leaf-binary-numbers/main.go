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

func sumRootToLeaf(root *TreeNode) int {
	sum := dfs(root, []int{})
	return sum
}

func dfs(root *TreeNode, nums []int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		var sum int
		for i := range nums {
			sum += nums[i] * int(math.Pow(2, float64(len(nums)-1-i)))
		}
		return sum
	}
	nums = append(nums, root.Val)
	return dfs(root.Left, nums) + dfs(root.Right, nums)
}

func createBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	tree := &TreeNode{nums[0], nil, nil}
	queue := []*TreeNode{tree}
	for i := 1; i < len(nums); i += 2 {
		cur := queue[0]
		queue = queue[1:]

		cur.Left = &TreeNode{nums[i], nil, nil}
		queue = append(queue, cur.Left)

		if i+1 < len(nums) {
			cur.Right = &TreeNode{nums[i+1], nil, nil}
			queue = append(queue, cur.Right)
		}
	}
	return tree
}

//iterative
func sumRootToLeaf2(root *TreeNode) int {
	type node struct {
		tree *TreeNode
		sum  int
	}
	stack := []node{node{root, 0}}
	var sum int

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curTree := cur.tree
		pathSum := cur.sum

		pathSum = pathSum << 1
		pathSum += curTree.Val

		if curTree.Left == nil && curTree.Right == nil {
			sum += pathSum
		}
		if curTree.Left != nil {
			stack = append(stack, node{curTree.Left, pathSum})
		}
		if curTree.Right != nil {
			stack = append(stack, node{curTree.Right, pathSum})
		}
	}
	return sum
}

func main() {
	root := []int{1, 0, 1, 0, 1, 0, 1}
	tree := createBinaryTree(root)
	fmt.Println(sumRootToLeaf2(tree))
}
