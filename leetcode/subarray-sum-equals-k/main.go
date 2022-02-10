package main

import (
	"fmt"
)

//hashmap
func subarraySum3(nums []int, k int) int {
	count, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := m[sum-k]; ok {
			count += m[sum-k]
		}
		m[sum]++
	}
	return count
}

//cumulative sum
func subarraySum2(nums []int, k int) int {
	sums := make([]int, len(nums)+1)

	sums[0] = 0
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sums[j]-sums[i] == k {
				count++
			}
		}
	}
	return count
}

//brute force
func subarraySum(nums []int, k int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			res := sum(nums[i:j])
			if res == k {
				ret++
			}
		}
	}
	return ret
}

func sum(nums []int) int {
	var ret int
	for _, v := range nums {
		ret += v
	}
	return ret
}

func main() {
	nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	k := 7
	fmt.Println(subarraySum3(nums, k))
}
