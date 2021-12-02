/*
https://leetcode.com/problems/same-tree/
Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:
	Input: p = [1,2,3], q = [1,2,3]
	Output: true

Example 2:
	Input: p = [1,2], q = [1,null,2]
	Output: false

Example 3:
	Input: p = [1,2,1], q = [1,1,2]
	Output: false
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var ret bool
	if (p != nil && q == nil) || (p == nil && q != nil){
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		ret = p.Val == q.Val
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return ret && left && right
}

func main() {

	a1 := &TreeNode{Val: 1}
	b1 := &TreeNode{Val: 2}
	c1 := &TreeNode{Val: 3, Left: b1, Right: a1}

	a2 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 2}
	c2 := &TreeNode{Val: 3, Left: b2, Right: a2}

    fmt.Println(isSameTree(c1, c2))
}
