package main

import "fmt"

func removeDuplicates(nums []int) int {

	var ret []int
	m := make(map[int]bool)
	for _, v := range nums {
		if !m[v] {
			m[v] = true
			ret = append(ret, v)
		}
	}
	copy(nums, ret)
	return len(ret)
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
