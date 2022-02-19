package main

import (
	"fmt"
)

//brute force
func maximumEvenSplit(finalSum int64) []int64 {

	output, _ := backtrack(finalSum, 2, 0, 0, make(map[int]bool), []int{})
	var ret []int64
	for i := range output {
		ret = append(ret, int64(output[i]))
	}
	return ret
}

func backtrack(finalSum int64, startNum, curSum, maxLen int, nums map[int]bool, output []int) ([]int, int) {
	if curSum == int(finalSum) && maxLen < len(nums) {
		var ret []int
		for k, _ := range nums {
			ret = append(ret, k)
		}
		maxLen = len(nums)
		return ret, len(nums)
	}

	for i := startNum; i <= int(finalSum); i += 2 {
		if curSum-i < int(finalSum) && !nums[i] {
			nums[i] = true
			curSum += i
			output, maxLen = backtrack(finalSum, startNum+2, curSum, maxLen, nums, output)
			delete(nums, i)
			curSum -= i
		}
	}
	return output, maxLen
}

func main() {
	fmt.Println(maximumEvenSplit(12))
	fmt.Println(maximumEvenSplit(7))
	fmt.Println(maximumEvenSplit(28))
	fmt.Println(maximumEvenSplit(90))
}
