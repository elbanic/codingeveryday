package main

import "fmt"

func checkString(s string) bool {
	if len(s) == 0 {
		return false
	}

	var foundB bool
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			foundB = true
		}
		if foundB {
			if s[i] == 'a' {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(checkString("aaabbb"))
	fmt.Println(checkString("abab"))
	fmt.Println(checkString("bbb"))
}
