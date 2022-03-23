/*
https://leetcode.com/problems/climbing-stairs/
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
	Input: n = 2
	Output: 2
	Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps

Example 2:
	Input: n = 3
	Output: 3
	Explanation: There are three ways to climb to the top.
	1. 1 step + 1 step + 1 step
	2. 1 step + 2 steps
	3. 2 steps + 1 step
*/

package main

import "fmt"

//brute force
//func climbStairs(n int) int {
//	return climbingCases(0, n)
//}
//
//func climbingCases(i, n int) int {
//	if i > n {
//		return 0
//	}
//	if i == n {
//		return 1
//	}
//	return climbingCases(i+1, n) + climbingCases(i+2, n)
//}

//memoizaion
func climbStairs(n int) int {
	memo := make(map[int]int)
	return climbingCases(0, n, memo)
}

func climbingCases(i, n int, m map[int]int) int {
	if v, ok := m[i]; ok {
		return v
	}
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	m[i] = climbingCases(i+1, n, m) + climbingCases(i+2, n, m)
	return m[i]
}

//tabulation
func climbStairs2(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)

	dp[0], dp[1] = 1, 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs2(35))
}
