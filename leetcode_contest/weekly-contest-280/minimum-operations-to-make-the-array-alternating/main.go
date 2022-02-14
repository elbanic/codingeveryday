package main

import (
	"fmt"
	"sort"
)

type keyCount struct {
	key       int
	occurence int
}

func minimumOperations(nums []int) int {

	even := make(map[int]int)
	odd := make(map[int]int)
	var maxEven, maxEvenCnt, maxOdd, maxOddCnt int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even[nums[i]]++
			if even[nums[i]] > maxEvenCnt {
				maxEvenCnt = even[nums[i]]
				maxEven = nums[i]
			}
		} else {
			odd[nums[i]]++
			if odd[nums[i]] > maxOddCnt {
				maxOddCnt = odd[nums[i]]
				maxOdd = nums[i]
			}
		}
	}

	if maxEven == maxOdd {
		var evenArr, oddArr []keyCount
		for k, v := range even {
			evenArr = append(evenArr, keyCount{k, v})
		}
		for k, v := range odd {
			oddArr = append(oddArr, keyCount{k, v})
		}
		sort.Slice(evenArr, func(i, j int) bool {
			return evenArr[i].occurence > evenArr[j].occurence
		})
		sort.Slice(oddArr, func(i, j int) bool {
			return oddArr[i].occurence > oddArr[j].occurence
		})

		for maxOdd == maxEven && (len(evenArr) > 0 || len(oddArr) > 0) {
			if len(evenArr) == 0 {
				maxOdd = oddArr[0].key
			} else if len(oddArr) == 0 {
				maxEven = evenArr[0].key
			} else {
				if evenArr[0].occurence > oddArr[0].occurence {
					maxEven = evenArr[0].key
				} else {
					maxOdd = oddArr[0].key
				}
				evenArr = evenArr[1:]
				oddArr = oddArr[1:]
			}
		}
	}

	if maxEven == maxOdd {
		maxOdd = maxEven + 1
	}

	var count int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			if nums[i] != maxEven {
				count++
			}
		} else {
			if nums[i] != maxOdd {
				count++
			}
		}
	}
	return count
}

func main() {
	nums := []int{2, 2}
	fmt.Println(minimumOperations(nums))

	nums2 := []int{1, 2, 2, 2, 2}
	fmt.Println(minimumOperations(nums2))

	nums3 := []int{4, 12, 91, 17, 29, 2, 32, 49, 5, 30, 89, 21, 91, 34, 71, 22, 88, 32, 36, 64, 28, 69, 7, 100, 35, 41, 62, 91, 85, 61, 4, 79, 77, 52, 57, 97, 41, 91, 13, 34, 37, 84, 10, 10, 37, 93, 58, 35, 81, 36, 81, 6, 50, 27, 68}
	fmt.Println(minimumOperations(nums3))
}
