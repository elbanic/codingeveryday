package main

import (
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)

func brokenCalc(startValue int, target int) int {

	var step int

	for target > startValue {
		step++
		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
	}
	return step + startValue - target
}

func main() {

	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))

}
