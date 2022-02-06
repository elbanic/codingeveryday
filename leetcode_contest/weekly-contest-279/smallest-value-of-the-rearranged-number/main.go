package main

import (
	"fmt"
	"math"
	"sort"
)

func smallestNumber(num int64) int64 {

	if num == 0 {
		return 0
	}

	minus := int64(1)
	if num < 0 {
		minus = -1
		num = num * minus
	}

	var digits []int64
	for num >= 1 {
		digits = append(digits, num%10)
		num /= 10
	}
	if minus == -1 {
		sort.Slice(digits, func(i, j int) bool {
			return digits[i] > digits[j]
		})
	} else {
		sort.Slice(digits, func(i, j int) bool {
			return digits[i] < digits[j]
		})
		if digits[0] == 0 {
			for i := 1; i < len(digits); i++ {
				if digits[i] != 0 {
					temp := digits[i]
					digits = append(digits[:i], digits[i+1:]...)
					digits = append(digits[:1], digits[0:]...)
					digits[0] = temp
					break
				}
			}
		}
	}

	var ret int64
	for i := 0; i < len(digits); i++ {
		ret += digits[i] * int64(math.Pow(float64(10), float64(len(digits)-1-i)))
	}
	return ret * minus
}

func main() {
	fmt.Println(smallestNumber(1384))
	fmt.Println(smallestNumber(-7605))
}
