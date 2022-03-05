package main

import (
	"fmt"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {

	m := make(map[int]int)
	for i := range mapping {
		m[i] = mapping[i]
	}

	type mapped struct {
		key int
		val int
	}
	var newNums []mapped
	for _, num := range nums {
		var newNum []int
		curNum := num
		if num == 0 {
			newNum = append(newNum, curNum%10)
		} else {
			for curNum > 0 {
				newNum = append(newNum, curNum%10)
				curNum /= 10
			}
		}

		var val int
		for i := len(newNum) - 1; i >= 0; i-- {
			val *= 10
			val += m[newNum[i]]
		}
		newNums = append(newNums, mapped{num, val})
	}
	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i].val < newNums[j].val
	})

	var ret []int
	for _, newNum := range newNums {
		ret = append(ret, newNum.key)
	}
	return ret
}

func main() {
	mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sortJumbled(mapping, nums))
}
