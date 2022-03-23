/*
https://leetcode.com/problems/word-pattern/
Given a pattern and a string s, find if s follows the same pattern.
Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:
	Input: pattern = "abba", s = "dog cat cat dog"
	Output: true

Example 2:
	Input: pattern = "abba", s = "dog cat cat fish"
	Output: false

Example 3:
	Input: pattern = "aaaa", s = "dog cat cat dog"
	Output: false
*/
package main

import (
	"fmt"
	"strings"
)

//single index hash map
func wordPattern2(pattern string, s string) bool {
	np := len(pattern)
	splitted := strings.Split(s, " ")
	ns := len(splitted)
	if np != ns {
		return false
	}
	m := make(map[string]int)
	for i := 0; i < ns; i++ {
		if _, ok := m["w_"+splitted[i]]; !ok {
			m["w_"+splitted[i]] = i
		}
		if _, ok := m["c_"+string(pattern[i])]; !ok {
			m["c_"+string(pattern[i])] = i
		}
		if m["w_"+splitted[i]] != m["c_"+string(pattern[i])] {
			return false
		}
	}
	return true
}

func wordPattern(pattern string, s string) bool {

	np := len(pattern)
	splitted := strings.Split(s, " ")
	ns := len(splitted)
	if np != ns {
		return false
	}

	m := make(map[string]string)
	for i := 0; i < np; i++ {
		if val, ok := m[string(pattern[i])]; ok {
			if val != splitted[i] {
				return false
			}
		} else {
			m[string(pattern[i])] = splitted[i]
		}
	}

	m = make(map[string]string)
	for i := 0; i < np; i++ {
		if val, ok := m[splitted[i]]; ok {
			if val != string(pattern[i]) {
				return false
			}
		} else {
			m[splitted[i]] = string(pattern[i])
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern2(pattern, s))

	pattern2 := "abba"
	s2 := "dog cat cat fish"
	fmt.Println(wordPattern2(pattern2, s2))

	pattern3 := "abba"
	s3 := "dog dog dog dog"
	fmt.Println(wordPattern2(pattern3, s3))

	pattern4 := "abc"
	s4 := "b c a"
	fmt.Println(wordPattern2(pattern4, s4))
}
