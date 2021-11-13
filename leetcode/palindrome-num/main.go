/*
https://leetcode.com/problems/palindrome-number/
Given an integer `x`, return `true` if `x` is palindrome integer.
An integer is a **palindrome** when it reads the same backward as forward. For example, `121` is palindrome while `123` is not.

Example 1:
	Input: x = 121
	Output: true

Example 2:
	Input: x = -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
	Input: x = 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Example 4:
	Input: x = -101
	Output: false

Constraints:
	-231 <= x <= 231 - 1

Approach
	I'm going to make a string array. Items in the array are input number converted to string
	and the numbers will be stored array in the order as the number of digits.
	The array size is the number of digits.
	The function will compare the first index and the end index of the array,
	and then will compare the second index and the second index to the last of the array.
*/

package main

import (
	"fmt"
)

//converting integer to string
func getSliceFromInt(x int) []int {

	s := make([]int, 0)
	i, iter := 0, x
	for {
		if iter == 0 {
			break
		}
		s = append(s, iter%10)
		iter /= 10
		i++
	}
	return s
}

//reverting the right half numbers
func revertedNumber(x int, r int) int {
	return (r * 10) + (x % 10)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x/10 < 0 {
		return true
	}
	if x % 10 == 0 && x != 0 {
		return false
	}

	//s := getSliceFromInt(x)
	//for i := 0; i < len(s)/2; i++ {
	//	if s[i] != s[len(s)-1-i] {
	//		return false
	//	}
	//}

	r := 0
	for x > r {
		r = revertedNumber(x, r)
		x /= 10
	}
	return x == r || x == r/10
}



func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
}
