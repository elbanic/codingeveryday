package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v] += 1
	}
	return m
}

func main() {
	q := "my name is my"
	m := WordCount(q)
	fmt.Println(m)
}
