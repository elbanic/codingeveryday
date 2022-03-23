/*
https://leetcode.com/problems/valid-perfect-square/
Given a positive integer num, write a function which returns True if num is a perfect square else False.
Follow up: Do not use any built-in library function such as sqrt.

Example 1:
	Input: num = 16
	Output: true

Example 2:
	Input: num = 14
	Output: false
*/
package main

import "fmt"

func isPerfectSquare(num int) bool {

	left := 0
	right := num

	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(left, mid, right)
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}
