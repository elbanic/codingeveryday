package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {

	s1Map := make(map[string]int)
	wind := make(map[string]int)

	for _, s := range s1 {
		s1Map[string(s)]++
	}
	wSize := len(s1)
	for i := 0; i < len(s2); i++ {
		if i < wSize {
			wind[string(s2[i])]++
		} else {
			if isSameMap(s1Map, wind) {
				return true
			}
			if v, ok := wind[string(s2[i-wSize])]; ok {
				if v <= 1 {
					delete(wind, string(s2[i-wSize]))
				} else {
					wind[string(s2[i-wSize])]--
				}
			}
			wind[string(s2[i])]++
		}
	}
	return isSameMap(s1Map, wind)
}

func isSameMap(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	s1 := "adc"
	s2 := "dcda"
	fmt.Println(checkInclusion(s1, s2))
}
