package main

import (
	"fmt"
	"math"
)

//tabulation
func minCostClimbingStairs(cost []int) int {

	if len(cost) == 0 || len(cost) == 1 {
		return 0
	}

	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(cost)+1; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+cost[i-1]), float64(dp[i-2]+cost[i-2])))
	}
	return dp[len(cost)]
}

//recursion
func minCostClimbingStairs2(cost []int) int {
	res := helper(len(cost), cost, make(map[int]int))
	return res
}

func helper(i int, cost []int, memo map[int]int) int {
	if i == 0 || i == 1 {
		return 0
	}
	if v, ok := memo[i]; ok {
		return v
	}
	first := helper(i-1, cost, memo) + cost[i-1]
	second := helper(i-2, cost, memo) + cost[i-2]
	res := int(math.Min(float64(first), float64(second)))
	memo[i] = res
	return res
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs2(cost))

	cost2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs2(cost2))
}
