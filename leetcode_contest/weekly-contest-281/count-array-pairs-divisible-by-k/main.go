package main

import (
	"fmt"
)

//brute force
func coutPairs2(nums []int, k int) int64 {
	var cnt int64
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]*nums[j])%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

func coutPairs(nums []int, k int) int64 {

	var div, unDiv []int
	for i := range nums {
		if nums[i]%k == 0 {
			div = append(div, nums[i])
		} else {
			unDiv = append(unDiv, nums[i])
		}
	}
	var combCnt int
	for i := len(div) - 1; i > 0; i-- {
		combCnt += i
	}

	firstCnt := len(div)*len(unDiv) + combCnt
	var lastCnt int
	for i := 0; i < len(unDiv)-1; i++ {
		for j := i + 1; j < len(unDiv); j++ {
			if (unDiv[i]*unDiv[j])%k == 0 {
				lastCnt++
			}
		}
	}

	return int64(firstCnt + lastCnt)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println(coutPairs(nums, k))
}
