package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//greedy with stack
func removeKdigits2(num string, k int) string {

	var nums []int
	for _, s := range num {
		val, _ := strconv.Atoi(string(s))
		nums = append(nums, val)
	}

	var stack []int
	for _, n := range nums {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > n {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, n)
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	var ret string
	for i := 0; i < len(stack); i++ {
		ret += strconv.Itoa(stack[i])
	}
	ret = strings.TrimLeft(ret, "0")
	if len(ret) == 0 {
		ret = "0"
	}
	return ret
}

//backtracking - Time Limit Exceeded
func removeKdigits(num string, k int) string {

	init, _ := strconv.Atoi(num[:len(num)-k])
	min := backtrack(num, "", 0, len(num)-k, init)

	return strconv.Itoa(min)
}

func backtrack(num string, prev string, index, length, min int) int {

	if len(prev) == length {
		res, _ := strconv.Atoi(prev)
		return int(math.Min(float64(res), float64(min)))
	}

	for i := index; i < len(num); i++ {
		prev += string(num[i])
		min = backtrack(num, prev, i+1, length, min)
		prev = prev[:len(prev)-1]
	}
	return min
}

func main() {
	num := "1432219"
	k := 3
	fmt.Println(removeKdigits2(num, k))

	num2 := "9"
	k2 := 1
	fmt.Println(removeKdigits2(num2, k2))
}
