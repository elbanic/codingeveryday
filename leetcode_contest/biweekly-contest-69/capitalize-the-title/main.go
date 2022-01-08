package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	strs := strings.Split(title, " ")
	var ret string
	for i, str := range strs {
		if len(str) <= 2 {
			if i != 0 {
				ret += " "
			}
			ret += strings.ToLower(str)
		} else {
			if i != 0 {
				ret += " "
			}
			first := string(str[0])
			ret += strings.ToUpper(first) + strings.ToLower(str[1:])
		}
	}
	return ret
}

func main() {
	fmt.Println(capitalizeTitle("capiTalIze tHe titLe"))
	fmt.Println(capitalizeTitle("First leTTeR of EACH Word"))
	fmt.Println(capitalizeTitle("i lOve leetcode"))
}
