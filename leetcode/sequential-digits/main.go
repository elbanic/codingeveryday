package main

import (
	"fmt"
	"math"
	"strconv"
)

//sliding window
func sequentialDigits2(low int, high int) []int {
	s := "123456789"
	lowDigit := int(math.Log10(float64(low))) + 1
	highDigit := int(math.Log10(float64(high))) + 1

	var ret []int
	for i := lowDigit; i <= highDigit; i++ {
		for j := 0; j < 10-i; j++ {
			cur, err := strconv.Atoi(s[j : j+i])
			if err == nil && cur >= low && cur <= high {
				ret = append(ret, cur)
			}
		}
	}
	return ret
}

func sequentialDigits(low int, high int) []int {

	digit := int(math.Log10(float64(low)))
	i := low / int(math.Pow(10, float64(digit)))

	var ret []int
	cur := low
	for cur < high {
		cur = genDigits(digit+1, i)
		if i+digit < 10 && cur >= low && cur <= high {
			ret = append(ret, cur)
		}
		if i == 9 {
			digit++
			i = 1
		} else {
			i++
		}
	}
	return ret
}

func genDigits(digit, i int) int {
	ret := i
	for j := 1; j < digit; j++ {
		cur := ret % 10
		ret *= 10
		ret += cur + 1
	}
	return ret
}

func main() {
	fmt.Println(sequentialDigits2(10, 1000000000))
	fmt.Println(sequentialDigits2(1000, 13000))
}
