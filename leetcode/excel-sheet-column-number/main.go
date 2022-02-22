package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {

	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := make(map[string]int)

	for i, c := range s {
		m[string(c)] = i + 1
	}

	n := len(s)
	var ret int
	//for i, c := range columnTitle {
	//	ret += m[string(c)] * int(math.Pow(float64(n), float64(len(columnTitle)-1-i)))
	//}
	for _, c := range columnTitle {
		ret *= n
		ret += m[string(c)]
	}
	return ret
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("ZZZ"))
}
