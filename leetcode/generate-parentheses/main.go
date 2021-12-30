/*
https://leetcode.com/problems/generate-parentheses/
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:
	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
	Input: n = 1
	Output: ["()"]
*/
package main

import "fmt"

//recursion
func generateParenthesis(n int) []string {
	output := generateParenthesisRec(n, 0, 0, "", []string{})
	return output
}

func generateParenthesisRec(n, open, close int, cur string, output []string) []string {
	if open == n && open == close {
		output = append(output, cur)
		return output
	}

	if open == n {
		output = generateParenthesisRec(n, open, close+1, cur+")", output)
	} else if open < n && open > close {
		output = generateParenthesisRec(n, open, close+1, cur+")", output)
		output = generateParenthesisRec(n, open+1, close, cur+"(", output)
	} else {
		output = generateParenthesisRec(n, open+1, close, cur+"(", output)
	}
	return output
}

//iteration
type parenthesisForm struct {
	str string
	op  int
	cl  int
}

func generateParenthesis2(n int) []string {

	stack := []parenthesisForm{parenthesisForm{"(", 1, 0}}

	var output []string
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.op == n && cur.op == cur.cl {
			output = append(output, cur.str)
		} else if cur.op == n {
			cur.str += ")"
			cur.cl++
			stack = append(stack, cur)
		} else if cur.op < n && cur.op > cur.cl {
			tmp := parenthesisForm{cur.str, cur.op, cur.cl}
			tmp.str += ")"
			tmp.cl++
			stack = append(stack, tmp)

			tmp2 := parenthesisForm{cur.str, cur.op, cur.cl}
			tmp2.str += "("
			tmp2.op++
			stack = append(stack, tmp2)
		} else {
			cur.str += "("
			cur.op++
			stack = append(stack, cur)
		}
	}
	return output
}

func main() {
	fmt.Println(generateParenthesis2(3))
}
