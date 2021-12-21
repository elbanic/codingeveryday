package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	var prefix []string

	for _, v := range strs[0] {
		prefix = append(prefix, string(v))
	}

	for i := 1; i < len(strs); i++ {
		if len(prefix) == 0 {
			break
		}
		j := 0
		for j < len(strs[i]) {
			if len(prefix)-1 < j {
				break
			}
			if prefix[j] != string(strs[i][j]) {
				break
			}
			if len(strs[i])-1 == j {
				prefix = prefix[:j+1]
			}
			j++
		}
		prefix = prefix[:j]
	}

	ret := strings.Join(prefix, "")
	return ret
}

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
