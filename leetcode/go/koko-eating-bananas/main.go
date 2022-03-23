package main

import (
	"fmt"
	"math"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {

	if len(piles) == 0 {
		return 0
	}
	if len(piles) == 1 {
		return int(math.Ceil(float64(piles[0]) / float64(h)))
	}

	sort.Ints(piles)
	left, right := 1, piles[len(piles)-1]

	for left < right {
		mid := left + (right-left)/2
		cur := calcTotalHours(piles, mid)
		if cur > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func calcTotalHours(piles []int, k int) int {
	totalH := 0
	for _, v := range piles {
		totalH += int(math.Ceil(float64(v) / float64(k)))
	}
	return totalH
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(piles, h))

	piles2 := []int{30, 11, 23, 4, 20}
	h2 := 5
	fmt.Println(minEatingSpeed(piles2, h2))

	piles3 := []int{30, 11, 23, 4, 20}
	h3 := 6
	fmt.Println(minEatingSpeed(piles3, h3))

	piles4 := []int{312884470}
	h4 := 312884469
	fmt.Println(minEatingSpeed(piles4, h4))
}
