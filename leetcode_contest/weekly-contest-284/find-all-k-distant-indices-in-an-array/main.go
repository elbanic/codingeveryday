package main

import (
	"fmt"
	"sort"
)

func findKDistantIndices(nums []int, key int, k int) []int {

	keyIndices := make(map[int]bool)
	for i, n := range nums {
		if n == key {
			keyIndices[i] = true
		}
	}

	retSet := make(map[int]bool)
	var ret []int
	for id := range keyIndices {
		for i := 0; i <= k; i++ {
			if id+i < len(nums) {
				if !retSet[id+i] {
					ret = append(ret, id+i)
				}
				retSet[id+i] = true
			}
			if id-i >= 0 {
				if !retSet[id-i] {
					ret = append(ret, id-i)
				}
				retSet[id-i] = true
			}
		}
	}
	sort.Ints(ret)
	return ret
}

func main() {
	nums := []int{3, 4, 9, 1, 3, 9, 5}
	key := 9
	k := 1

	fmt.Println(findKDistantIndices(nums, key, k))
}
