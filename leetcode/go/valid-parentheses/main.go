/*
https://leetcode.com/problems/valid-parentheses/
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:
	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.

Example 1:
	Input: s = "()"
	Output: true

Example 2:
	Input: s = "()[]{}"
	Output: true

Example 3:
	Input: s = "(]"
	Output: false

Example 4:
	Input: s = "([)]"
	Output: false

Example 5:
	Input: s = "{[]}"
	Output: true
*/

package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	pair := make(map[string]string)
	pair["("] = ")"
	pair["["] = "]"
	pair["{"] = "}"

	var stack []string
	for i := 0; i < len(s); i++ {
		if _, ok := pair[string(s[i])]; ok {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if string(s[i]) == pair[top] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(("))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
