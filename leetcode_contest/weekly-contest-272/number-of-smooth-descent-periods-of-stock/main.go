package main

import "fmt"

//brute force
func getDescentPeriods(prices []int) int64 {
	start := 0
	end := 0
	var descents [][]int
	for i := 0; i < len(prices); i++ {
		start = i
		end = i
		descents = append(descents, prices[start:start+1])
		for end+1 < len(prices) {
			if prices[end]-1 != prices[end+1] {
				break
			}
			end++
			descents = append(descents, prices[start:end])
		}
	}
	return int64(len(descents))
}

func main() {
	prices := []int{3, 2, 1, 4}
	fmt.Println(getDescentPeriods(prices))
}
