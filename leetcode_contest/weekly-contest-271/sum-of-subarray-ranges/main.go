package main

import (
	"fmt"
	"sync"
)

func subArrayRanges(nums []int) int64 {

	divSize := 2
	if len(nums)/divSize < 1 {
		var sum int
		for subSize := 2; subSize < len(nums)+1; subSize++ {
			for i := 0; i < len(nums)-subSize+1; i++ {
				subArray := nums[i : i+subSize]
				max := Max(subArray)
				min := Min(subArray)
				sum += max - min
			}
		}
		return int64(sum)
	} else {
		var wg sync.WaitGroup
		var subArrs [][]int
		result := make(map[int]int)
		for i := 0; i < len(nums)+1; i += divSize {
			subArrs = append(subArrs, nums[i:i+divSize])
		}

		for i, v := range subArrs {
			wg.Add(1)
			go func(slice []int, result map[int]int, id int) {
				var sum int
				for subSize := 2; subSize < len(slice)+1; subSize++ {
					for i := 0; i < len(slice)-subSize+1; i++ {
						subArray := slice[i : i+subSize]
						max := Max(subArray)
						min := Min(subArray)
						sum += max - min
					}
				}
				wg.Done()
				result[id] = sum
			}(v, result, i)
		}
		wg.Wait()
	}
	return 0
}

func Max(nums []int) int {
	var max int
	for i, e := range nums {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}

func Min(nums []int) int {
	var min int
	for i, e := range nums {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}

func main() {
	//nums := []int{1, 2, 3}
	//fmt.Println(subArrayRanges(nums))

	nums2 := []int{4, -2, -3, 4, 1}
	fmt.Println(subArrayRanges(nums2))
}
