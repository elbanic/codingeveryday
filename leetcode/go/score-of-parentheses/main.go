package main

import (
	"fmt"
	"math"
)

func scoreOfParentheses(s string) int {

	stack := []int{0}
	for _, c := range s {
		if string(c) == "(" {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			w := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, w+int(math.Max(float64(2*v), float64(1))))
		}
	}
	return stack[len(stack)-1]
}

func main() {
	s := "(()(()))"
	fmt.Println(scoreOfParentheses(s))

	s1 := "(((())))"
	fmt.Println(scoreOfParentheses(s1))
}
