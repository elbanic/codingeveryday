/*
https://leetcode.com/problems/insert-into-a-binary-search-tree/
You are given the root node of a binary search tree (BST) and a value to insert into the tree.
Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.
Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion.
You can return any of them.

Example 1:
	Input: root = [4,2,7,1,3], val = 5
	Output: [4,2,7,1,3,5]

Example 2:
	Input: root = [40,20,60,10,30,50,70], val = 25
	Output: [40,20,60,10,30,50,70,null,null,25]

Example 3:
	Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
	Output: [4,2,7,1,3,5]
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//recursion
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{val, nil, nil}
		return root
	}
	if root.Val < val {
		root.Right = insertIntoBST2(root.Right, val)
	} else {
		root.Left = insertIntoBST2(root.Left, val)
	}
	return root
}

//iteration
func insertIntoBST(root *TreeNode, val int) *TreeNode {

	cur := root
	for cur != nil {
		if cur.Val < val {
			if cur.Right == nil {
				cur.Right = &TreeNode{val, nil, nil}
				return root
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Left == nil {
				cur.Left = &TreeNode{val, nil, nil}
				return root
			} else {
				cur = cur.Left
			}
		}
	}
	return &TreeNode{val, nil, nil}
}

func createBinarySearchTree(nums []int) *TreeNode {
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
		if i < len(nums)-1 {
			cur.Right = &TreeNode{nums[i+1], nil, nil}
			queue = append(queue, cur.Right)
		}
	}
	return tree
}

func printTree(tree *TreeNode) {
	if tree != nil {
		fmt.Printf("%d ", tree.Val)
		printTree(tree.Left)
		printTree(tree.Right)
	}
	return
}

func main() {
	root := []int{4, 2, 7, 1, 3}
	tree := createBinarySearchTree(root)
	val := 5
	printTree(tree)
	fmt.Println()
	tree = insertIntoBST2(tree, val)
	printTree(tree)
	fmt.Println()
}
