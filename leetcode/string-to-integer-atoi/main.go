package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {

	s = strings.Trim(s, " ")
	start := 0
	end := 0
	var found bool
	var signDetect bool

	for i := 0; i < len(s); i++ {
		_, err := strconv.Atoi(string(s[i]))
		if !found && err == nil {
			start = i
			found = true
		} else if found && err != nil {
			end = i
			break
		}
		if string(s[i]) == "-" || string(s[i]) == "+" {
			if signDetect {
				return 0
			} else {
				signDetect = true
			}
		} else if !found {
			return 0
		}
	}

	if found && end <= start {
		end = len(s)
	}
	ret, _ := strconv.Atoi(s[start:end])
	if start > 0 {
		if s[start-1] == '-' {
			ret = -ret
		}
	}
	min := int(math.Pow(-2, 31))
	max := int(math.Pow(2, 31) - 1)
	if ret < min {
		return min
	} else if ret > max {
		return max
	}
	return ret
}

func main() {
	fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("        -42"))
	fmt.Println(myAtoi("985 words"))
	fmt.Println(myAtoi("words 985"))
}
