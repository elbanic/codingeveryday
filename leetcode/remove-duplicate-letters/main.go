package main

import (
	"fmt"
	"strings"
)

func removeDuplicateLetters(s string) string {

	m := make(map[string][]int)
	for i, c := range s {
		if _, ok := m[string(c)]; ok {
			m[string(c)] = append(m[string(c)], i)
		} else {
			m[string(c)] = []int{i}
		}
	}

	var stack []string
	visited := make(map[string]bool)
	for id, c := range s {
		if !visited[string(c)] {

			for len(stack) > 0 {
				pop := stack[len(stack)-1]
				if pop > string(c) {
					arr, _ := m[pop]
					if len(arr) > 0 {
						removeId := findIndex(arr, id)
						if removeId == -1 {
							m[pop] = []int{}
						} else {
							m[pop] = m[pop][removeId+1:]
							stack = stack[:len(stack)-1]
							delete(visited, pop)
						}
					} else {
						break
					}
				} else {
					break
				}
			}
			stack = append(stack, string(c))
			visited[string(c)] = true
		}
	}
	return strings.Join(stack, "")
}

func findIndex(nums []int, id int) int {
	for i := range nums {
		if nums[i] > id {
			return i - 1
		}
	}
	return -1
}

func main() {
	s := "abacb"
	fmt.Println(removeDuplicateLetters(s))
}
