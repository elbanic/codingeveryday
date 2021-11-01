/*
https://leetcode.com/problems/palindrome-number/
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
	"math"
)

func getLength(x int) int {
	size := 0
	iter := x
	for {
		if iter = iter / 10; iter != 0 {
			size += 1
		} else {
			break
		}
	}
	if size != 0 {
		size += 1
	}
	return size
}

func isPalindrome(x int) bool {
	if x < 0 || (-231 > x) || (x > 230) {
		return false
	}
	len := getLength(x)
	if len < 2 {
		return false
	}

	s := make([]int, len)
	j := 0
	iter := x
	for i := math.Pow(10., float64(len-1)); i >= 1; i /= 10 {
		if iter != 0 {
			s[j] = iter / int(i)
			iter %= int(i)
			j++
		}
	}

	for i := 0; i < len/2; i++ {
		if s[i] != s[len-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
}
