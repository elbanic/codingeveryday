package main

import (
	"fmt"
	"math"
)

func divideString(s string, k int, fill byte) []string {

	var ret []string

	i := 0
	for i < len(s)-k {
		endId := int(math.Min(float64(i+k), float64(len(s))))
		ret = append(ret, s[i:endId])
		i += k
	}
	if i < len(s) {
		temp := s[i:len(s)]
		for j := 0; j < k-(len(s)-i); j++ {
			temp += string(fill)
		}
		ret = append(ret, temp)
	}

	return ret
}

func main() {
	s := "abcdefghij"
	k := 3
	fmt.Println(divideString(s, k, 'x'))
}
