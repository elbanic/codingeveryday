package main

import "fmt"

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
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
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(words))
}
