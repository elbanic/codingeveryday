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

//maxDepth for top-down
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	depth, maxDepth := 1, 1
//	findMaxDepth(root, depth, &maxDepth)
//	return maxDepth
//}
//
//func findMaxDepth(root *TreeNode, depth int, maxDepth *int) {
//
//	if root == nil {
//		return
//	}
//	if root.Left == nil && root.Right == nil {
//		*maxDepth = int(math.Max(float64(*maxDepth), float64(depth)))
//	}
//
//	findMaxDepth(root.Left, depth + 1, maxDepth)
//	findMaxDepth(root.Right, depth + 1, maxDepth)
//}

//maxDepth for bottom-up
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findMaxDepth(root)
}

func findMaxDepth(root *TreeNode) int{

	if root == nil {
		return 0
	}
	left := findMaxDepth(root.Left)
	right := findMaxDepth(root.Right)
	return int(math.Max(float64(left), float64(right))) + 1
}

func main() {
	left := TreeNode{Val: 3}
	right := TreeNode{Val: 2, Left: &left}
	root := TreeNode{Val: 1, Left: nil, Right: &right}
	fmt.Println(preorderTraversal(&root))
	fmt.Println(inorderTraversal(&root))
	fmt.Println(postorderTraversal(&root))

	A := TreeNode{Val: 5}
	B := TreeNode{Val: 4}
	C := TreeNode{Val: 3}
	D := TreeNode{Val: 2, Left: &B, Right: &A}
	E := TreeNode{Val: 1, Left: &D, Right: &C}

	fmt.Println(levelOrder(&E))

	F := TreeNode{Val: 2}
	G := TreeNode{Val: 1, Left: &F}
	fmt.Println(levelOrder(&G))

	a := TreeNode{Val: 7}
	b := TreeNode{Val: 15}
	c := TreeNode{Val: 20, Left: &b, Right: &a}
	d := TreeNode{Val: 9}
	e := TreeNode{Val: 3, Left: &d, Right: &c}
	fmt.Println(maxDepth(&e))
}
