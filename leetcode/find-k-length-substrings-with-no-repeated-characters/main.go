package main

import "fmt"

func numKLenSubstrNoRepeats(s string, k int) int {

	if len(s) < k {
		return 0
	}

	var count int
	m := make(map[string]int)
	for i := 0; i < k; i++ {
		m[string(s[i])]++
	}
	if len(m) == k {
		count++
	}

	prev := 0
	for i := k; i < len(s); i++ {
		if v, ok := m[string(s[prev])]; ok {
			if v == 1 {
				delete(m, string(s[prev]))
			} else {
				m[string(s[prev])]--
			}
		}
		m[string(s[i])]++
		if len(m) == k {
			fmt.Println(m)
			count++
		}
		prev++
	}
	return count
}

func main() {
	s := "abab"
	k := 2
	fmt.Println(numKLenSubstrNoRepeats(s, k))
}
