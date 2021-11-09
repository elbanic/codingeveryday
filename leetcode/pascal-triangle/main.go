/*
https://leetcode.com/problems/pascals-triangle/
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:
	Input: numRows = 5
	Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example 2:
	Input: numRows = 1
	Output: [[1]]

Constraints:
	1 <= numRows <= 30

*/

package main

import "fmt"

func generateTriangle(depth int, prev []int, triangle [][]int) [][]int {
	triangle = append(triangle, prev)
	depth--
	if depth == 0 {
		return triangle
	}
	var cur []int
	for i := 0; i < len(prev); i++ {
		if i == 0 {
			cur = append(cur, prev[i])
		}
		j := i + 1
		if j != len(prev) {
			cur = append(cur, prev[i]+prev[j])
		}
		if i == len(prev)-1 {
			cur = append(cur, 1)
		}
	}
	return generateTriangle(depth, cur, triangle)
}

func generate(numRows int) [][]int {
	var triangle [][]int
	seed := []int{1}
	return generateTriangle(numRows, seed, triangle)
}

func main() {

	fmt.Println(generate(5))
	fmt.Println(generate(1))

}
