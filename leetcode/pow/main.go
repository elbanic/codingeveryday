/*
https://leetcode.com/problems/powx-n/
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:

	Input: x = 2.00000, n = 10
	Output: 1024.00000

Example 2:
	Input: x = 2.10000, n = 3
	Output: 9.26100

Example 3:
	Input: x = 2.00000, n = -2
	Output: 0.25000
	Explanation: 2-2 = 1/22 = 1/4 = 0.25
 */
package main

import "fmt"

//divide and conquer + recursive
func helperPow(x float64, n int) float64{
	if n == 0 {
		return 1
	}

	half := helperPow(x, n / 2)

	if n % 2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return helperPow(x, n)
}

func main() {
	a2 := myPow(2, 3)
	a1 := myPow(2, 10)
	a3 := myPow(2, -2)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
