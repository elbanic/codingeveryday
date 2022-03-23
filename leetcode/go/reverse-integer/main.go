/*
https://leetcode.com/problems/reverse-integer/
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
	Input: x = 123
	Output: 321

Example 2:
	Input: x = -123
	Output: -321

Example 3:
	Input: x = 120
	Output: 21

Example 4:
	Input: x = 0
	Output: 0
*/
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	minus := false
	if x < 0 {
		minus = true
		x = -x
	}
	var nums []int
	for x >= 1 {
		digit := x % 10
		nums = append(nums, digit)
		x = x / 10
	}
	var ret int
	for i, v := range nums {
		ret += v * int(math.Pow(10, float64(len(nums)-1-i)))
	}
	if ret > int(math.Pow(2, 31)-1) {
		return 0
	}

	if minus {
		return -ret
	}
	return ret
}

//Pop and Push Digits & Check before Overflow
func reversePopPush(x int) int {
	var rev int

	for x != 0 {
		pop := x % 10
		x /= 10
		//Integer MAX
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev * 10 + pop
	}
	return rev
}

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reversePopPush(-123))
}
