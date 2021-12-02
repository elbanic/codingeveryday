package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	index := new(int)
	*index = -1
	succ := helperInorderSuccessors(root, p, []*TreeNode{}, index)

	if *index == -1 || len(succ)-1 < *index+1{
		return nil
	}
	return succ[*index + 1]
}

func helperInorderSuccessors(root *TreeNode, p *TreeNode, succ []*TreeNode, index *int) []*TreeNode {
	if root == nil {
		return succ
	}
	succ = helperInorderSuccessors(root.Left, p, succ, index)
	if root.Val == p.Val {
		*index = len(succ)
	}
	succ = append(succ, root)
	succ = helperInorderSuccessors(root.Right, p, succ, index)
	return succ
}

func inorderTraversal(root *TreeNode){
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Println(root.Val)
	inorderTraversal(root.Right)
}

func main() {
	a := &TreeNode{Val: 3}
	b := &TreeNode{Val: 1}
	c := &TreeNode{Val: 2, Left: b, Right: a}
	inorderTraversal(c)

	a1 := &TreeNode{Val: 6}
	b1 := &TreeNode{Val: 1}
	c1 := &TreeNode{Val: 2, Left: b1}
	d1 := &TreeNode{Val: 4}
	e1 := &TreeNode{Val: 3, Left: c1, Right: d1}
	f1 := &TreeNode{Val: 5, Left: e1, Right: a1}
	result := inorderSuccessor(f1, a1)
	if result != nil {
		fmt.Println(result.Val)
	} else {
		fmt.Println("nil")
	}

	a2 := &TreeNode{Val: 3}
	b2 := &TreeNode{Val: 2, Right: a2}
	result2 := inorderSuccessor(b2, b2)
	if result2 != nil {
		fmt.Println(result2.Val)
	} else {
		fmt.Println("nil")
	}

}
