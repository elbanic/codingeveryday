/*
https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b
where v = |a.val - b.val| and a is an ancestor of b.
A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.

Example 1:
	Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
	Output: 7
	Explanation: We have various ancestor-node differences, some of which are given below :
	|8 - 3| = 5
	|3 - 7| = 4
	|8 - 1| = 7
	|10 - 13| = 3
	Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.

*/
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

//brute force
func maxAncestorDiff(root *TreeNode) int {
	max := inorderTraversal(root, []int{}, -1)
	return max
}

func inorderTraversal(node *TreeNode, parent []int, max int) int {
	if node == nil {
		return max
	}

	if node.Left == nil && node.Right == nil {
		parent = append(parent, node.Val)
		for i := 0; i < len(parent); i++ {
			for j := i; j < len(parent); j++ {
				max = int(math.Max(math.Abs(float64(parent[i]-parent[j])), float64(max)))
			}
		}
		return max
	}
	parent = append(parent, node.Val)
	max = inorderTraversal(node.Left, parent, max)
	max = inorderTraversal(node.Right, parent, max)
	return max
}

func createBTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := TreeNode{nums[0], nil, nil}
	queue := []*TreeNode{&root}
	for i := 1; i < len(nums); i += 2 {
		cur := queue[0]
		queue = queue[1:]
		if nums[i] == -1 {
			cur.Left = nil
		} else {
			left := &TreeNode{nums[i], nil, nil}
			cur.Left = left
			queue = append(queue, left)
		}
		if i+1 < len(nums) {
			if nums[i+1] == -1 {
				cur.Right = nil
			} else {
				right := &TreeNode{nums[i+1], nil, nil}
				cur.Right = right
				queue = append(queue, right)
			}
		}
	}
	return &root
}

func main() {
	root := []int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13}
	tree := createBTree(root)
	fmt.Println(maxAncestorDiff(tree))
}
