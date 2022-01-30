package main

import (
	"fmt"
	"math"
)

//brute force
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {

	alph := "abcdefghijklmnopqrstuvwxyz"
	alphaVal := make(map[string]int)
	for i, v := range alph {
		alphaVal[string(v)] = i + 1
	}

	for j := 0; j <= len(s)-k; j++ {
		if hashValue == hash(alphaVal, s[j:j+k], power, modulo) {
			return s[j : j+k]
		}
	}
	return ""
}

func hash(val map[string]int, s string, p int, m int) int {
	var ret float64
	for i := 0; i < len(s); i++ {
		mul := math.Mod(math.Pow(float64(p), float64(i)), float64(m))
		ret += float64(val[string(s[i])]) * mul
		ret = math.Mod(ret, float64(m))
	}
	return int(ret)
}

func main() {
	fmt.Println(subStrHash("leetcode", 7, 20, 2, 0))
	fmt.Println(subStrHash("fbxzaad", 31, 100, 3, 32))
	fmt.Println(subStrHash("xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk", 22, 51, 41, 9))
}
