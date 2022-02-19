package main

import "fmt"

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			j, k := i+1, 1
			for j < i+len(needle) && j < len(haystack) {
				if haystack[j] != needle[k] {
					break
				}
				j++
				k++
			}
			if j == i+len(needle) {
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))
}
