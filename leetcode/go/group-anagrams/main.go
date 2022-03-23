package main

import (
	"fmt"
	"sort"
)

// categorize by sorted string
func groupAnagrams2(strs []string) [][]string {
	ret := make(map[string][]string)
	for _, str := range strs {
		chars := toCharSet(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := toString(chars)
		if _, ok := ret[key]; !ok {
			ret[key] = []string{}
		}
		ret[key] = append(ret[key], str)
	}
	var group [][]string
	for _, v := range ret {
		group = append(group, v)
	}
	return group
}

func toCharSet(str string) []rune {
	var ret []rune
	for _, v := range str {
		ret = append(ret, v)
	}
	return ret
}

func toString(chars []rune) string {
	var ret string
	for _, v := range chars {
		ret += string(v)
	}
	return ret
}

func groupAnagrams(strs []string) [][]string {
	var group [][]string
	for len(strs) > 0 {
		str := strs[0]
		strs = strs[1:]

		newGroup := []string{str}
		dict := make(map[string]int)
		var idx []int
		for i := range str {
			dict[string(str[i])]++
		}
		for i, other := range strs {

			var found bool
			if len(str) == len(other) {
				newDict := make(map[string]int)
				for k, v := range dict {
					newDict[k] = v
				}
				for _, ch := range other {
					_, ok := newDict[string(ch)]
					if ok {
						newDict[string(ch)]--
						if newDict[string(ch)] == 0 {
							delete(newDict, string(ch))
						}
					} else {
						found = false
						break
					}
				}
				if len(newDict) == 0 {
					found = true
				}
			}
			if found {
				newGroup = append(newGroup, other)
				idx = append(idx, i)
			}
		}
		group = append(group, newGroup)
		if len(idx) > 0 {
			for i := len(idx) - 1; i >= 0; i-- {
				strs = append(strs[:idx[i]], strs[idx[i]+1:]...)
			}
		}
	}
	return group
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams2(strs))
}
