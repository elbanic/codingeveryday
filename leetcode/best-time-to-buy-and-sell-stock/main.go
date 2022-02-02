package main

import (
	"fmt"
	"math"
)

func maxProfit3(prices []int) int {
	min := int(^uint(0) >> 1)
	var maxprifit int

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > maxprifit {
			maxprifit = prices[i] - min
		}
	}
	return maxprifit
}

//next max - time limit exceeded
func maxProfit2(prices []int) int {

	start := detectInc(prices)
	if start == -1 {
		return 0
	}

	max := maxVal(prices[start:])
	var ret int
	for i := start; i < len(prices); i++ {
		if prices[i] != max {
			if max-prices[i] > ret {
				ret = max - prices[i]
			}
		} else {
			max = maxVal(prices[i+1:])
		}
	}
	return ret
}

func maxVal(nums []int) int {
	var max int
	for _, v := range nums {
		max = int(math.Max(float64(v), float64(max)))
	}
	return max
}

func detectInc(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return i
		}
	}
	return -1
}

//brute force - time limit exceeded
func maxProfit(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	var max int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[i] < prices[j] && prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit3(prices))

	prices2 := []int{7, 5, 8, 4, 1, 5}
	fmt.Println(maxProfit3(prices2))
}
