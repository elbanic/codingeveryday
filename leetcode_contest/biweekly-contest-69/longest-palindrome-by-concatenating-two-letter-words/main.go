package main

import "fmt"

func longestPalindrome(words []string) int {
	jokers := make(map[string]int)
	var candidates []string

	var jokerSum int
	var i int
	for len(words) > 0 {
		found := false
		if words[i][0] == words[i][1] {
			jokers[words[i]]++
			jokerSum += 2
			words = append(words[:i], words[i+1:]...)
		}
		for j := i + 1; j < len(words); j++ {
			if isPalindrome(words[i] + words[j]) {
				candidates = append(candidates, words[i])
				candidates = append(candidates, words[j])
				words = append(words[:j], words[j+1:]...)
				words = append(words[:i], words[i+1:]...)
				found = true
				break
			}
		}
		if !found && len(words) > 0 {
			words = append(words[:i], words[i+1:]...)
		}
	}
	var ret int
	var jokerMax int
	for _, v := range jokers {
		if jokerMax < v {
			jokerMax = v
		}
	}
	if jokerSum > 0 && (jokerSum/2)%2 == 0 {
		jokerSum = (jokerSum/2 - 1) * 2
	}

	ret += len(candidates) * 2
	ret += jokerSum
	return ret
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if string(s[i]) != string(s[len(s)-1-i]) {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"ll", "lb", "bb", "bx", "xx", "lx", "xx", "lx", "ll", "xb", "bx", "lb", "bb", "lb", "bl", "bb", "bx", "xl", "lb", "xx"}
	fmt.Println(longestPalindrome(words))
}
