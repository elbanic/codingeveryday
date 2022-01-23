package main

import (
	"fmt"
)

func rearrangeArray(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}

	var ret []int
	var pos, neg []int

	for _, v := range nums {
		if v > 0 {
			pos = append(pos, v)
		} else {
			neg = append(neg, v)
		}
	}
	for i := 0; i < len(nums)/2; i++ {
		ret = append(ret, pos[i])
		ret = append(ret, neg[i])
	}
	return ret
}

func main() {
	nums := []int{-1, 1}
	fmt.Println(rearrangeArray(nums))
}
