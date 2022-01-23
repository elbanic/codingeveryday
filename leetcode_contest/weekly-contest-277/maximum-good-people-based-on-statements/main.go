package main

import "fmt"

func maximumGood(statements [][]int) int {

	var maxBad, maxBadId int
	for i := 0; i < len(statements); i++ {
		var badCount int
		for j := 0; j < len(statements); j++ {
			if statements[i][j] == 0 {
				badCount++
			}
		}
		if maxBad < badCount {
			maxBad = badCount
			maxBadId = i
		}
	}
	if maxBad > 0 {
		for i := range statements[maxBadId] {
			if statements[maxBadId][i] == 0 {
				statements[maxBadId][i] = 1
			}
		}
	}

	var ret int
	m := make(map[int]bool)
	for i := 0; i < len(statements); i++ {
		for j := 0; j < len(statements); j++ {
			if statements[i][j] == 1 {
				if !m[j] {
					m[j] = true
					ret++
				}
			}
		}
	}

	return ret
}

func main() {
	statements := [][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}
	fmt.Println(maximumGood(statements))
}
