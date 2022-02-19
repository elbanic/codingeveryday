package main

import "fmt"

//brute force
func countPairs(nums []int, k int) int {

	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (i*j)%k == 0 && nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	fmt.Println(countPairs(nums, k))
}
