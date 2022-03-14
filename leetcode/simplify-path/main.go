package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {

	parr := strings.Split(path, "/")
	var stack []string

	for _, p := range parr {
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if p != "." && p != "" {
			stack = append(stack, p)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/../"
	fmt.Println(simplifyPath(path))
}
