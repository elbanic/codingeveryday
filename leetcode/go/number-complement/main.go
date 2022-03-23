/*
https://leetcode.com/problems/number-complement/
The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer num, return its complement.

Example 1:
	Input: num = 5
	Output: 2
	Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.

Example 2:
	Input: num = 1
	Output: 0
	Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
*/
package main

import (
	"fmt"
	"math"
)

func findComplement(num int) int {

	bin := toBin(num)
	for i := range bin {
		if bin[i] == 1 {
			bin[i] = 0
		} else {
			bin[i] = 1
		}
	}
	return toInt(bin)
}

func toBin(num int) []int {
	var ret []int
	for num > 1 {
		digit := num % 2
		ret = append(ret, digit)
		num = num / 2
	}
	return ret
}

func toInt(bins []int) int {
	var ret int
	for i := range bins {
		ret += bins[i] * int(math.Pow(2, float64(i)))
	}
	return ret
}

func main() {
	fmt.Println(findComplement(1))
}
