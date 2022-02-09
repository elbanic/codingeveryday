package main

import (
	"fmt"
	"math"
	"sort"
)

func findPairs2(nums []int, k int) int {

	sort.Ints(nums)
	var ret int
	if k == 0 {
		m := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			m[nums[i]]++
		}
		for _, v := range m {
			if v > 1 {
				ret++
			}
		}
	} else {
		m := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			m[nums[i]] = true
		}

		nums = deDup(nums)
		for i := 0; i < len(nums); i++ {
			if m[nums[i]+k] {
				ret++
			}
		}
	}
	return ret
}

func deDup(nums []int) []int {
	m := make(map[int]bool)
	for _, i := range nums {
		m[i] = true
	}
	var ret []int
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

type pair struct {
	x int
	y int
}

func findPairs(nums []int, k int) int {

	m := make(map[pair]bool)
	var ret int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			max := int(math.Max(float64(nums[i]), float64(nums[j])))
			min := int(math.Min(float64(nums[i]), float64(nums[j])))
			_, ok := m[pair{min, max}]

			if !ok {
				diff := int(math.Abs(float64(nums[i] - nums[j])))
				m[pair{min, max}] = true
				if diff == k {
					ret++
				}
			}
		}
	}
	return ret
}

func main() {
	nums := []int{1, 3, 1, 5, 4}
	k := 0
	fmt.Println(findPairs2(nums, k))
}
