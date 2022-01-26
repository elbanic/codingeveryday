package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//iteration
func getAllElements2(root1 *TreeNode, root2 *TreeNode) []int {

	var stack1, stack2 []*TreeNode
	curr1 := root1
	curr2 := root2

	var ret []int
	for curr1 != nil || len(stack1) > 0 || curr2 != nil || len(stack2) > 0 {

		for curr1 != nil {
			stack1 = append(stack1, curr1)
			curr1 = curr1.Left
		}
		for curr2 != nil {
			stack2 = append(stack2, curr2)
			curr2 = curr2.Left
		}
		if len(stack2) == 0 || len(stack1) > 0 && stack1[len(stack1)-1].Val <= stack2[len(stack2)-1].Val {
			curr1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]

			ret = append(ret, curr1.Val)
			curr1 = curr1.Right
		} else {
			curr2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]

			ret = append(ret, curr2.Val)
			curr2 = curr2.Right
		}
	}
	return ret
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	r1arr := inOrderTraverse(root1, []int{})
	r2arr := inOrderTraverse(root2, []int{})

	r1arr = append(r1arr, r2arr...)
	sort.Ints(r1arr)
	return r1arr
}

func inOrderTraverse(root *TreeNode, output []int) []int {
	if root == nil {
		return output
	}
	if root.Left != nil {
		output = inOrderTraverse(root.Left, output)
	}
	output = append(output, root.Val)
	if root.Right != nil {
		output = inOrderTraverse(root.Right, output)
	}
	return output
}

func createTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return &TreeNode{}
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
	root1 := createTree([]int{2, 1, 4})
	root2 := createTree([]int{1, 0, 3})

	fmt.Println(getAllElements2(root1, root2))

	root12 := createTree([]int{1, -1, 8})
	root22 := createTree([]int{8, 1})

	fmt.Println(getAllElements2(root12, root22))
}
