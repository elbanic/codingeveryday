package main

import "fmt"

func prefixCount(words []string, pref string) int {

	var cnt int
	for _, word := range words {
		if len(word) >= len(pref) {
			isSamePref := true
			for j := 0; j < len(pref); j++ {
				if word[j] != pref[j] {
					isSamePref = false
				}
			}
			if isSamePref {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	pref := "at"
	fmt.Println(prefixCount(words, pref))
}
