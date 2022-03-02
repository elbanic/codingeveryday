package main

import "fmt"

func isSubsequence(s string, t string) bool {

	var i, j int
	nS := len(s)
	nT := len(t)
	for i < nS && j < nT {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == nS
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s1 := "axc"
	t1 := "ahbgdc"
	fmt.Println(isSubsequence(s1, t1))

	s2 := "aec"
	t2 := "abcde"
	fmt.Println(isSubsequence(s2, t2))
}
