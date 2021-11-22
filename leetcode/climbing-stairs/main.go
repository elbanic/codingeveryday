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

//cases
//func climbStairs(n int) int {
//	output := climbingCases(1, n, []int{}, [][]int{})
//	for _,v := range climbingCases(2, n, []int{}, [][]int{}) {
//		output = append(output, v)
//	}
//	return len(output)
//}
//
//func climbingCases(x, n int, path []int, output [][]int) [][]int {
//
//	path = append(path, x)
//	if n - x == 0{
//		output = append(output, path)
//		return output
//	} else if n - x < 0 {
//		return output
//	}
//	n -= x
//	if n == 1 {
//		output = climbingCases(1, n, path, output)
//	} else {
//		output = climbingCases(1, n, path, output)
//		output = climbingCases(2, n, path, output)
//	}
//	return output
//}

func main() {
	fmt.Println(climbStairs(35))
}
