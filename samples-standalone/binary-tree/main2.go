package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int{
	var ret []int
	preorderT(root, &ret)
	return ret
}

func preorderT(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorderT(root.Left, output)
		preorderT(root.Right, output)
	}
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	inorderT(root, &ret)
	return ret
}

func inorderT(root *TreeNode, output *[]int) {
	if root != nil {
		inorderT(root.Left, output)
		*output = append(*output, root.Val)
		inorderT(root.Right, output)
	}
}

func postorderTraversal(root *TreeNode) []int {
	var ret []int
	postorderT(root, &ret)
	return ret
}

func postorderT(root *TreeNode, output *[]int) {
	if root != nil {
		postorderT(root.Left, output)
		postorderT(root.Right, output)
		*output = append(*output, root.Val)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ret [][]int
	ret = append(ret, []int{root.Val})

	var queue []*TreeNode
	if root.Left != nil {
		queue = append(queue, root.Left)
	}
	if root.Right != nil {
		queue = append(queue, root.Right)
	}

	for len(queue) > 0 {
		var currVals []int
		size := len(queue)

		for i:=0; i<size; i++{
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			currVals = append(currVals, curr.Val)
		}
		ret = append(ret, currVals)
	}
	return ret
}

func main() {
	left := TreeNode{Val: 3}
	right := TreeNode{Val: 2, Left: &left}
	root := TreeNode{Val: 1, Left: nil, Right: &right}
	fmt.Println(preorderTraversal(&root))
	fmt.Println(inorderTraversal(&root))
	fmt.Println(postorderTraversal(&root))

	//[3,9,20,null,null,15,7]
	A := TreeNode{Val: 5}
	B := TreeNode{Val: 4}
	C := TreeNode{Val: 3}
	D := TreeNode{Val: 2, Left: &B, Right: &A}
	E := TreeNode{Val: 1, Left: &D, Right: &C}

	fmt.Println(levelOrder(&E))

	F := TreeNode{Val: 2}
	G := TreeNode{Val: 1, Left: &F}

	fmt.Println(levelOrder(&G))
}
