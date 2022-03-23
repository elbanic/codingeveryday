package main

import "fmt"

func findTheDifference(s string, t string) byte {

	m := make(map[string]int)
	for _, i := range s {
		m[string(i)]++
	}

	for _, j := range t {
		if val, ok := m[string(j)]; ok {
			if val > 1 {
				m[string(j)]--
			} else {
				delete(m, string(j))
			}
		} else {
			return byte(j)
		}
	}
	return 0
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(findTheDifference(s, t))
}
