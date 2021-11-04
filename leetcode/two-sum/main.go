package main

import (
	"fmt"
	"sync"
)

func twoSum(nums []int, target int) []int {

	ch := make(chan []int)
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(nums []int) {
		wg.Done()
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					v := []int{i, j}
					ch <- v
				}
			}
		}
	}(nums)
	wg.Wait()

	results := <-ch
	return results
}

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
	twoSum([]int{3, 2, 4}, 6)
	fmt.Println("")
}
