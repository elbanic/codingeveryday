/*
https://leetcode.com/problems/decode-string/
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
For example, there won't be input like 3a or 2[4].

Example 1:
	Input: s = "3[a]2[bc]"
	Output: "aaabcbc"

Example 2:
	Input: s = "3[a2[c]]"
	Output: "accaccacc"

Example 3:
	Input: s = "2[abc]3[cd]ef"
	Output: "abcabccdcdcdef"

Example 4:
	Input: s = "abc3[cd]xyz"
	Output: "abccdcdcdxyz"
*/
package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return decode(s, 1)
}

func decode(s string, r int) string {
	var ret string
	for r > 0 {
		for i := 0; i < len(s); i++ {
			if val, err := strconv.Atoi(string(s[i])); err == nil {
				//update the digit
				start := i + 1
				for start < len(s) {
					if string(s[start]) == "[" {
						val, _ = strconv.Atoi(s[i:start])
						break
					}
					start++
				}

				opened := 0
				end := start
				for end < len(s) {
					if string(s[end]) == "[" {
						opened++
					} else if string(s[end]) == "]" {
						opened--
					}
					if opened == 0 {
						break
					}
					end++
				}
				ret += decode(s[start+1:end], val)
				i = end
			} else {
				ret += string(s[i])
			}
		}
		r--
	}
	return ret
}

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("100[leetcode]"))
}
