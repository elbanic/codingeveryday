package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {

	var stack []int
	for len(pushed) > 0 || len(popped) > 0 {
		if len(stack) == 0 {
			if len(pushed) > 0 {
				stack = append(stack, pushed[0])
				pushed = pushed[1:]
			}
		} else {
			top := stack[len(stack)-1]
			if len(popped) > 0 && top == popped[0] {
				stack = stack[:len(stack)-1]
				popped = popped[1:]
			} else if len(pushed) > 0 {
				stack = append(stack, pushed[0])
				pushed = pushed[1:]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))

	pushed2 := []int{1, 2, 3, 4, 5}
	popped2 := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed2, popped2))
}
