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

//recursion
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil{
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right
}

//iteration
func isSameTreeIteration(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queueP := []*TreeNode{p}
	queueQ := []*TreeNode{q}

	for len(queueP) > 0 {
		p = queueP[len(queueP)-1]
		q = queueQ[len(queueQ)-1]

		queueP = queueP[0:len(queueP)-1]
		queueQ = queueQ[0:len(queueQ)-1]

		if p.Val != q.Val {
			return false
		}
		if (p.Left != nil && q.Left == nil) || (p.Left == nil && q.Left != nil){
			return false
		}
		if p.Left != nil && q.Left != nil{
			queueP = append(queueP, p.Left)
			queueQ = append(queueQ, q.Left)
		}
		if (p.Right != nil && q.Right == nil) || (p.Right == nil && q.Right != nil){
			return false
		}
		if p.Right != nil && q.Right != nil{
			queueP = append(queueP, p.Right)
			queueQ = append(queueQ, q.Right)
		}
	}

	return true
}

func main() {

	a1 := &TreeNode{Val: 1}
	b1 := &TreeNode{Val: 2}
	c1 := &TreeNode{Val: 3, Left: b1, Right: a1}

	a2 := &TreeNode{Val: 1}
	b2 := &TreeNode{Val: 2}
	c2 := &TreeNode{Val: 3, Left: b2, Right: a2}
    fmt.Println(isSameTreeIteration(c1, c2))

	d1 := &TreeNode{}
	d2 := &TreeNode{}
	fmt.Println(isSameTreeIteration(d1, d2))
}
