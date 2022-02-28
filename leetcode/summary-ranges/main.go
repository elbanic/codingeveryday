package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var ret []string
	var start, cur int
	for cur < len(nums) {
		if cur != len(nums)-1 && nums[cur]+1 < nums[cur+1] {
			ret = append(ret, getRangeString(nums[start], nums[cur]))
			start = cur + 1
		}
		cur++
	}
	ret = append(ret, getRangeString(nums[start], nums[len(nums)-1]))
	return ret
}

func getRangeString(n1, n2 int) string {
	if n1 == n2 {
		return strconv.Itoa(n1)
	} else {
		return strconv.Itoa(n1) + "->" + strconv.Itoa(n2)
	}
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
