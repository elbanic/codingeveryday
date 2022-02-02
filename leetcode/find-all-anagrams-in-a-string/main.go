package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	windSize := len(p)
	pMap := make(map[string]int)
	for _, v := range p {
		pMap[string(v)]++
	}

	curMap := make(map[string]int)
	for i := 0; i < windSize; i++ {
		curMap[string(s[i])]++
	}

	var ret []int
	if isSameMap(pMap, curMap) {
		ret = append(ret, 0)
	}

	for i := 0; i < len(s)-windSize; i++ {
		if val, ok := curMap[string(s[i])]; ok {
			if val > 1 {
				curMap[string(s[i])]--
			} else {
				delete(curMap, string(s[i]))
			}
		}
		curMap[string(s[i+windSize])]++
		if isSameMap(pMap, curMap) {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func isSameMap(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for key, val := range m1 {
		if m2[key] != val {
			return false
		}
	}
	return true
}

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}
