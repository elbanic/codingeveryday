package main

import (
	"fmt"
	"strings"
)

func detectCapitalUse(word string) bool {

	letter := "abcdefghijklmnopqrstuvwxyz"
	letterM := make(map[string]bool)
	letterLM := make(map[string]bool)

	for _, s := range letter {
		letterM[string(s)] = true
		letterLM[strings.ToUpper(string(s))] = true
	}

	isStartCapital := letterLM[string(word[0])]
	allSmall, allCapital := true, true
	for i := 1; i < len(word); i++ {
		allSmall = allSmall && letterM[string(word[i])]
		allCapital = allCapital && letterLM[string(word[i])]
	}

	if isStartCapital {
		return allSmall || allCapital
	}
	return allSmall
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FlaG"))
	fmt.Println(detectCapitalUse("leetcode"))
}
