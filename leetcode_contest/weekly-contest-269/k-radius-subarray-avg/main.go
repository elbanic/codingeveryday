package main

import "fmt"

func getAverages(nums []int, k int) []int {
	ret := make([]int, len(nums))

	for i:=0; i<len(nums); i++ {
		if i < k || i >= len(nums) - k{
			ret[i] = -1
		} else {
			sub := nums[i-k:i+k+1]
			ret[i] = avg(sub, len(sub))
		}
	}
	return ret
}

func avg(nums []int, k int) int {
	var sum int
	for _,v := range nums {
		sum += v
	}
	return sum/k
}

func main() {
	nums := []int{7,4,3,9,1,8,5,2,6}
	k := 3

	fmt.Println(getAverages(nums,k))

}
