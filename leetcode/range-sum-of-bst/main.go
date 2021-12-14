/*
https://leetcode.com/problems/range-sum-of-bst/
Given the root node of a binary search tree and two integers low and high,
return the sum of values of all nodes with a value in the inclusive range [low, high].

Example 1:
	Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
	Output: 32
	Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

Example 2:
	Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
	Output: 23
	Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var sum int
	return traverseSum(root, low, high, sum)
}

func traverseSum(node *TreeNode, low int, high int, sum int) int {
	if node == nil {
		return sum
	}
	if node.Val >= low && node.Val <= high {
		sum += node.Val
	}
	sum = traverseSum(node.Left, low, high, sum)
	sum = traverseSum(node.Right, low, high, sum)
	return sum
}

func main() {

	f := TreeNode{18, nil, nil}
	e := TreeNode{7, nil, nil}
	d := TreeNode{3, nil, nil}
	c := TreeNode{15, nil, &f}
	b := TreeNode{5, &d, &e}
	a := TreeNode{10, &b, &c}
	fmt.Println(rangeSumBST(&a, 7, 15))
}
