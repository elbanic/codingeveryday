package main

import "fmt"

func splitSlice(s []int, size int) [][]int {
	var ret [][]int
	for i := 0; i < len(s); i += size {
		split := i + size

		if split > len(s) {
			split = len(s)
		}
		ret = append(ret, s[i:split])
	}
	return ret
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	split := splitSlice(s, 2)
	fmt.Println(split)
}
