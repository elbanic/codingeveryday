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

import (
	"fmt"
)

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

/*
https://leetcode.com/problems/pascals-triangle-ii/
Pascal's Triangle II
Given an integer `rowIndex`, return the `rowIndex`th (**0-indexed**) row of the **Pascal's triangle**.
In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:
Example 1:
	Input: rowIndex = 3
	Output: [1,3,3,1]

Example 2:
	Input: rowIndex = 0
	Output: [1]

Example 3:
	Input: rowIndex = 1
	Output: [1,1]
 */
func getRow(rowIndex int) []int {
	return getLastRowTriangle(0, rowIndex + 1, []int{})
}

func getLastRowTriangle(depth, size int, prev []int) []int{

	if depth == size {
		return prev
	}
	var curr []int
	for i:=0; i<=depth; i++ {
		if i == 0 || i == depth {
			curr = append(curr, 1)
		} else {
			curr = append(curr, prev[i-1] + prev[i])
		}
	}
	return getLastRowTriangle(depth+1, size, curr)
}

func recursiveTriangle(depth, size int, output [][]int) [][]int{

	if depth == size {
		return output
	}
	var row []int
	for i:=0; i<=depth; i++ {
		if i == 0 || i == depth {
			row = append(row, 1)
		} else {
			prevRow := output[len(output)-1]
			row = append(row, prevRow[i-1] + prevRow[i])
		}
	}
	output = append(output, row)
	return recursiveTriangle(depth+1, size, output)
}

func main() {

	fmt.Println(generate(5))
	fmt.Println(generate(1))

	fmt.Println(getRow(3))

	ret := recursiveTriangle(0, 5, [][]int{})
	fmt.Println(ret)
}
