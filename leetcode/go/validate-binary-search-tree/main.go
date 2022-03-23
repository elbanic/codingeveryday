/*
https://leetcode.com/problems/validate-binary-search-tree/
Given the root of a binary tree, determine if it is a valid binary search tree (BST).
A valid BST is defined as follows:
	The left subtree of a node contains only nodes with keys less than the node's key.
	The right subtree of a node contains only nodes with keys greater than the node's key.
	Both the left and right subtrees must also be binary search trees.

Example 1:
	Input: root = [2,1,3]
	Output: true

Example 2:
	Input: root = [5,1,4,null,null,3,6]
	Output: false
	Explanation: The root node's value is 5 but its right child's value is 4.
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *int

func isValidBST(root *TreeNode) bool {
	prev = nil
	return helperValidInorderBST(root)
}

//Recursive Inorder Traversal
func helperValidInorderBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !helperValidInorderBST(root.Left) {
		return false
	}
	if prev != nil && root.Val <= *prev {
		return false
	}
	if prev == nil {
		prev = new(int)
	}
	*prev = root.Val
	return helperValidInorderBST(root.Right)
}

//Recursive Traversal with Valid Range
func helperValidBST(left, right *int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if left != nil && node.Val <= *left ||
		right != nil && node.Val >= *right {
		return false
	}
	leftStatus := helperValidBST(left, &node.Val, node.Left)
	rightStatus := helperValidBST(&node.Val, right, node.Right)
	return leftStatus && rightStatus
}

func main() {
	b := &TreeNode{Val: 1}
	c := &TreeNode{Val: 1, Left: b, Right: nil}
	fmt.Println(isValidBST(c))
}
