package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for _,v := range words1 {
		m1[v] += 1
	}
	for _,v := range words2 {
		m2[v] += 1
	}

	var ret int
	for key,val1 := range m1 {
		val2 := m2[key]
		if val1 == 1 && val2 == 1{
			ret++
		}
	}
	return ret
}

func main() {
	words1 := []string{"leetcode","is","amazing","as","is"}
	words2 := []string{"amazing","leetcode","is"}
	fmt.Println(countWords(words1, words2))

	words3 := []string{"a","ab"}
	words4 := []string{"a","a","a","ab"}
	fmt.Println(countWords(words3, words4))
}
