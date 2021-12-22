/*

 */

package main

import "fmt"

func canConvert(str1 string, str2 string) bool {
	if str1 == "abcdefghijklmnopqrstuvwxyz" {
		return true
	}
	map1 := make(map[string][]int)
	map2 := make(map[string][]int)

	for i, s := range str1 {
		if arr, ok := map1[string(s)]; ok {
			arr = append(arr, i)
			map1[string(s)] = arr
		} else {
			map1[string(s)] = []int{i}
		}
	}
	for i, s := range str2 {
		if arr, ok := map2[string(s)]; ok {
			arr = append(arr, i)
			map2[string(s)] = arr
		} else {
			map2[string(s)] = []int{i}
		}
	}

	for _, v := range map1 {
		if !containsValue(map2, v) {
			return false
		}
	}
	return true
}

func containsValue(m map[string][]int, v []int) bool {
	for _, x := range m {
		if len(x) == len(v) {
			ret := true
			for i := range v {
				if v[i] != x[i] {
					ret = false
				}
			}
			if ret {
				return ret
			}
		}
	}
	return false
}

func canConvert2(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	conversionMappings := make(map[string]string)
	uniqueCharactersInStr2 := make(map[string]bool)

	for i := 0; i < len(str1); i++ {
		if v, ok := conversionMappings[string(str1[i])]; ok {
			if v != string(str2[i]) {
				return false
			}
		} else {
			conversionMappings[string(str1[i])] = string(str2[i])
			uniqueCharactersInStr2[string(str2[i])] = true
		}
	}
	if len(uniqueCharactersInStr2) < 26 {
		return true
	}
	return false
}

func main() {
	str1, str2 := "abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyzz"
	fmt.Println(canConvert(str1, str2))
}
