package main

import (
	"fmt"
)

func findStrobogrammatic(n int) []string {
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	lefts := left(n, 0, "", []string{})
	var output []string

	if n%2 == 1 {
		middle := []string{"0", "1", "8"}
		var temp []string
		for _, mid := range middle {
			for _, left := range lefts {
				temp = append(temp, left+mid)
			}
		}
		lefts = temp
	}

	for _, left := range lefts {
		stack := left
		if n%2 == 1 {
			stack = stack[:len(stack)-1]
		}
		output = append(output, right(stack, left))
	}
	return output
}

func left(n, i int, prev string, output []string) []string {
	if n/2 == i {
		output = append(output, prev)
		return output
	}
	var base []string
	if i == 0 {
		base = []string{"1", "6", "8", "9"}
	} else {
		base = []string{"0", "1", "6", "8", "9"}
	}

	for _, num := range base {
		prev += num
		output = left(n, i+1, prev, output)
		prev = prev[:len(prev)-1]
	}
	return output
}

func right(stack string, output string) string {
	if len(stack) == 0 {
		return output
	}
	last := string(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	switch last {
	case "6":
		output += "9"
	case "9":
		output += "6"
	default:
		output += last
	}
	return right(stack, output)
}

func main() {
	fmt.Println(findStrobogrammatic(5))
	fmt.Println(findStrobogrammatic(4))
	fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(1))
}
