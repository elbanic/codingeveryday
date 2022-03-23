package main

import "fmt"

//backtrack [Time Limit Exceeded]
func minRemoveToMakeValid(s string) string {
	var N int
	for _, c := range s {
		if string(c) == "(" || string(c) == ")" {
			N++
		}
	}
	for i := 0; i <= N; i++ {
		ret, ok := helper(s, i)
		if ok {
			return ret
		}
	}
	return ""
}

func helper(s string, n int) (string, bool) {
	if n == 0 {
		return s, isValidParentheses(s)
	}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == ")" {
			ret, ok := helper(s[:i]+s[i+1:], n-1)
			if ok {
				return ret, true
			}
		}
	}
	return s, false
}

func isValidParentheses(s string) bool {
	var stack []string
	for _, c := range s {
		if string(c) == "(" {
			stack = append(stack, string(c))
		} else if string(c) == ")" {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()))((ksg(((u((()c)("
	fmt.Println(minRemoveToMakeValid(s))
}
