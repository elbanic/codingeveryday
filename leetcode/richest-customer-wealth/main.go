package main

import "fmt"

func maximumWealth(accounts [][]int) int {

	if len(accounts) == 0 {
		return 0
	}

	var max int
	for _, account := range accounts {
		var sum int
		for _, money := range account {
			sum += money
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(accounts))

	accounts2 := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts2))
}
