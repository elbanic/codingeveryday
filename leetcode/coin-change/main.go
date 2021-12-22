/*
https://leetcode.com/problems/coin-change/
You are given an integer array coins representing coins of different denominations
and an integer amount representing a total amount of money.
Return the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an infinite number of each kind of coin.

Example 1:
	Input: coins = [1,2,5], amount = 11
	Output: 3
	Explanation: 11 = 5 + 5 + 1

Example 2:
	Input: coins = [2], amount = 3
	Output: -1

Example 3:
	Input: coins = [1], amount = 0
	Output: 0
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

//brute force - my solution
const MAX_INT = int(^uint(0) >> 1)

func coinChange(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	return recursionCoinChange(coins, amount, MAX_INT, 0)
}

func recursionCoinChange(coins []int, amount int, min int, depth int) int {
	if amount == 0 {
		return depth
	}
	if amount > 0 && depth < min {
		for i := 0; i < len(coins); i++ {
			curMin := recursionCoinChange(coins, amount-coins[i], min, depth+1)
			if curMin != -1 {
				if min > curMin {
					min = curMin
				}
			}
		}
		if min == MAX_INT {
			return -1
		} else {
			return min
		}
	}
	return -1
}

//brute force
func coinChange2(coins []int, amount int) int {
	return coinChangeHelper(0, coins, amount)
}
func coinChangeHelper(idxCoin int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if idxCoin < len(coins) && amount > 0 {
		maxVal := amount / coins[idxCoin]
		minCost := MAX_INT
		for x := 0; x <= maxVal; x++ {
			if amount >= x*coins[idxCoin] {
				res := coinChangeHelper(idxCoin+1, coins, amount-x*coins[idxCoin])
				if res != -1 {
					minCost = int(math.Min(float64(minCost), float64(res+x)))
				}
			}
		}
		if minCost == MAX_INT {
			return -1
		} else {
			return minCost
		}
	}
	return -1
}

//DP
func coinChangeDP(coins []int, amount int) int {

	if amount < 1 {
		return 0
	}
	count := make([]int, amount)
	return coinChangeDPHelper(coins, amount, count)
}

func coinChangeDPHelper(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if count[rem-1] != 0 {
		return count[rem-1]
	}

	min := MAX_INT
	for _, coin := range coins {
		res := coinChangeDPHelper(coins, rem-coin, count)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == MAX_INT {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}

func main() {
	//coins := []int{186, 419, 83, 408}
	//amount := 6249
	coins := []int{1, 2, 5}
	amount := 11
	//coins := []int{2, 3, 5}
	//amount := 17
	//fmt.Println(coinChange(coins, amount))

	fmt.Println(coinChangeDP(coins, amount))
}
