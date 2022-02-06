package main

import (
	"fmt"
)

type LeftRight struct {
	Left  int
	Right int
}

func minimumTime(s string) int {

	var comb []LeftRight
	min := int(^uint(0) >> 1)
	left := 0
	for left < len(s) {
		if string(s[left]) == "1" {
			right := 0
			for i := left + 1; i < len(s); i++ {
				if string(s[i]) == "1" {
					right = len(s) - i
					break
				}
			}
			lr := LeftRight{left + 1, right}
			comb = append(comb, lr)
			if min > lr.Left+lr.Right {
				min = lr.Left + lr.Right
			}
		}
		left++
	}

	right := len(s) - 1
	for right >= 0 {
		if string(s[right]) == "1" {
			left := -1
			for i := right - 1; i >= 0; i-- {
				if string(s[i]) == "1" {
					left = i
					break
				}
			}
			lr := LeftRight{left + 1, len(s) - right}
			comb = append(comb, lr)
			if min > lr.Left+lr.Right {
				min = lr.Left + lr.Right
			}
		}
		right--
	}

	return min
}

func main() {
	fmt.Println(minimumTime("011001111111101 0010100000010100 11"))
	fmt.Println(minimumTime("10001010110"))
}
