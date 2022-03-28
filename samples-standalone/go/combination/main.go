/*
Recursive Algorithm for combination array with given number
*/

package main

import "fmt"

func combination(depth int, slice []string, prev []string, comb [][]string) [][]string {
	if len(slice) == 0 || depth <= 0 {
		return comb
	}
	depth--
	for i := range slice {
		temp := make([]string, len(prev)+1)
		copy(temp, prev)
		temp[len(temp)-1] = slice[i]
		if depth == 0 {
			comb = append(comb, temp)
		}
		comb = combination(depth, slice[i+1:], temp, comb)
	}
	return comb
}

func main() {
	num := 3
	s := []string{"A", "B", "C", "D", "E"}
	var prev []string
	var comb [][]string
	comb = combination(num, s, prev, comb)

	for i := range comb {
		fmt.Println(comb[i])
	}
}
