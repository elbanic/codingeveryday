/*
https://leetcode.com/problems/palindrome-partitioning/
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.
A palindrome string is a string that reads the same backward as forward.

Example 1:
	Input: s = "aab"
	Output: [["a","a","b"],["aa","b"]]

Example 2:
	Input: s = "a"
	Output: [["a"]]
*/
package main

import "fmt"

//backtracking with DP
func partition3(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	var res [][]string
	res = dfs2(res, s, 0, []string{}, dp)
	return res
}

func dfs2(res [][]string, s string, start int, currentList []string, dp [][]bool) [][]string {
	if start >= len(s) {
		tmp := make([]string, len(currentList))
		copy(tmp, currentList)
		res = append(res, tmp)
		return res
	}
	for end := start; end < len(s); end++ {
		if s[start] == s[end] && (end-start <= 2 || dp[start+1][end-1]) {
			dp[start][end] = true
			currentList = append(currentList, s[start:end+1])
			res = dfs2(res, s, end+1, currentList, dp)
			currentList = currentList[:len(currentList)-1]
		}
	}
	return res
}

//backtracking
func partition2(s string) [][]string {
	var res [][]string
	res = dfs(0, res, []string{}, s)
	return res
}

func dfs(start int, res [][]string, currentList []string, s string) [][]string {
	if start >= len(s) {
		tmp := make([]string, len(currentList))
		copy(tmp, currentList)
		res = append(res, tmp)
		return res
	}
	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			currentList = append(currentList, s[start:end+1])
			res = dfs(end+1, res, currentList, s)
			currentList = currentList[:len(currentList)-1]
		}
	}
	return res
}

func isPalindrome(s string, low, high int) bool {
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}

func partition(s string) [][]string {

	var ret [][]string
	start, slide, count := 0, 0, 0
	var found bool

	var one []string
	for _, v := range s {
		one = append(one, string(v))
	}
	ret = append(ret, one)

	prev := "#"
	var partitions []string
	for slide < len(s) {
		if !found {
			count++
			if prev == string(s[slide]) {
				found = true
				count--
			}
		} else {
			count--
			if string(s[start+count]) != string(s[slide]) {
				found = false
			}
			if count == 0 {
				partitions = append(partitions, s[start:slide])
				start = slide + 1
			}
		}
		prev = string(s[slide])
		slide++
	}
	ret = append(ret, partitions)
	return ret
}

func main() {
	fmt.Println(partition3("cbbbcc"))
	fmt.Println(partition3("abccba"))
}
