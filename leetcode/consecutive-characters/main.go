/*
https://leetcode.com/problems/consecutive-characters/
The power of the string is the maximum length of a non-empty substring that contains only one unique character.
Given a string s, return the power of s.

Example 1:
	Input: s = "leetcode"
	Output: 2
	Explanation: The substring "ee" is of length 2 with the character 'e' only.

Example 2:
	Input: s = "abbcccddddeeeeedcba"
	Output: 5
	Explanation: The substring "eeeee" is of length 5 with the character 'e' only.

Example 3:
	Input: s = "triplepillooooow"
	Output: 5

Example 4:
	Input: s = "hooraaaaaaaaaaay"
	Output: 11

Example 5:
	Input: s = "tourist"
	Output: 1
*/

package main

import "fmt"

func maxPower(s string) int {

	if len(s) == 0 {
		return 0
	}

	cur := string(s[0])
	max := 1
	count := 1
	for i := 1; i < len(s); i++ {
		if string(s[i]) == cur {
			count++
			if count > max {
				max = count
			}
		} else {
			cur = string(s[i])
			count = 1
		}
	}
	return max
}

func main() {
	s := "tourist"
	fmt.Println(maxPower(s))
}
